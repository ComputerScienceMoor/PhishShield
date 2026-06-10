/**
 * URL Feature Extraction for Phishing Detection
 * Extracts 9 URL-only features matching the UCI dataset definitions
 */

/**
 * PhishShield – Enhanced URL-only features with Punycode / Homograph detection
 */

function isPunycodeDomain(hostname) {
  // Quick check: contains 'xn--' (punycode prefix)
  return hostname.includes("xn--");
}

function hasSuspiciousHomograph(hostname) {
  // Common homograph replacements (Latin look-alikes)
  const homographMap = {
    a: ["а", "ɑ"], // Cyrillic a, Greek alpha
    b: ["Ь", "Ƅ"], // Cyrillic soft sign variants
    c: ["с", "ϲ"], // Cyrillic c, Coptic c
    e: ["е", "ҽ"], // Cyrillic e
    i: ["і", "ӏ"], // Cyrillic i, lowercase L with stroke
    j: ["ј"], // Cyrillic je
    o: ["о", "ο", "օ"], // Cyrillic o, Greek omicron
    p: ["р"], // Cyrillic p
    x: ["х"], // Cyrillic x
    y: ["у", "γ"], // Cyrillic y, Greek gamma
    l: ["ӏ", "ł"], // Often confused with 'I' or '1'
    0: ["Ο", "О"], // Zero vs O
  };

  // Convert hostname to punycode-normalized form (if browser supports)
  let normalized = hostname;
  try {
    normalized = new URL(`http://${hostname}`).hostname; // forces punycode decoding
  } catch {}

  // Check for mixed scripts or look-alikes
  for (let i = 0; i < hostname.length; i++) {
    const char = hostname[i];
    const latin = Object.keys(homographMap).find((k) =>
      homographMap[k].includes(char)
    );
    if (latin && char !== latin) {
      // Found a homograph replacement
      return true;
    }
  }

  // Bonus: check for suspicious script mixing (Latin + Cyrillic/Greek)
  const hasCyrillic = /[\u0400-\u04FF]/.test(hostname);
  const hasGreek = /[\u0370-\u03FF]/.test(hostname);
  const hasLatin = /[a-zA-Z]/.test(hostname);

  return (hasCyrillic || hasGreek) && hasLatin; // mixed scripts = suspicious
}

/**
 * Extract all 9 URL features from a URL string
 * @param {string} urlString - The URL to analyze
 * @returns {Array<number>} Array of 9 feature values
 */
function extractURLFeatures(urlString) {
  try {
    let url = urlString.trim();

    if (!url.match(/^https?:\/\//i)) url = "http://" + url;

    const parsed = new URL(url);
    const domain = parsed.hostname;
    const href = parsed.href.toLowerCase();
    const pathname = parsed.pathname.toLowerCase();
    const search = parsed.search.toLowerCase();

    const urlLength = url.length;
    const domainLength = domain.length;
    const noOfSubDomain = (domain.match(/\./g) || []).length - 1;
    const isDomainIP = /^\d{1,3}(\.\d{1,3}){3}$/.test(domain) ? 1 : 0;
    const isHTTPS = url.toLowerCase().startsWith("https://") ? 1 : 0;
  
    const noOfDigits = (url.match(/\d/g) || []).length;
    const noOfOtherSpecial = (url.match(/[^a-zA-Z0-9.:/?=&\-_]/g) || []).length;
    const noOfEquals = (url.match(/=/g) || []).length;
    const noOfQMark = (url.match(/\?/g) || []).length;
    const noOfAmpersand = (url.match(/&/g) || []).length;
  
    const specialRatio = urlLength > 0 ? noOfOtherSpecial / urlLength : 0;
    const letterRatio =
      urlLength > 0 ? (url.match(/[a-zA-Z]/g) || []).length / urlLength : 0;
    const digitRatio = urlLength > 0 ? noOfDigits / urlLength : 0;
    const noOfLetters = (url.match(/[a-zA-Z]/g) || []).length;
    const tld = domain.split(".").pop() || "";
    const tldLength = tld.length;

    const features = [
      urlLength,
      domainLength,
      noOfSubDomain,
      isDomainIP,
      isHTTPS,
      noOfDigits,
      noOfOtherSpecial,
      noOfEquals,
      noOfQMark,
      noOfAmpersand,
      specialRatio,
      letterRatio,
      digitRatio,
      noOfLetters,
      tldLength,
    ];

    // ─── Additional lightweight checks (do NOT change array size) ────────
    // Suspicious TLDs (common in 2024–2025 phishing)
    const suspiciousTLDs = [
      "top",
      "xyz",
      "online",
      "site",
      "club",
      "shop",
      "store",
      "live",
      "digital",
      "buzz",
      "fun",
      "monster",
      "loan",
      "win",
      "men",
      "gq",
      "cf",
      "ml",
      "ga",
      "tk",
    ];
    
    const tldSuspicious = suspiciousTLDs.includes(tld);

    // Suspicious keywords in domain / path / query
    const scamKeywords = [
      "login",
      "signin",
      "account",
      "verify",
      "update",
      "secure",
      "bank",
      "paypal",
      "apple",
      "microsoft",
      "amazon",
      "recovery",
      "reset",
      "bursary",
      "scholarship",
      "claim",
      "prize",
      "congrat",
      "free",
      "gift",
    ];
    const keywordMatch = scamKeywords.some(
      (kw) =>
        domain.includes(kw) || pathname.includes(kw) || search.includes(kw)
    );

    // Boost phishing probability if either trigger is present
    if (tldSuspicious || keywordMatch) {
      features[0] = 100000; 
      features[3] = 1
      features[4] = 0; 
      features[14] = 50
    }

    const isPuny = isPunycodeDomain(domain);
    const isHomograph = hasSuspiciousHomograph(domain);

    if (isPuny || isHomograph) {
      // Boost phishing signals in existing features
      features[0] = 100000; 
      features[3] = 1; // treat as suspicious IP-like
      features[4] = 0; // force prefix-suffix suspicious
    }

    if (Shortining_Service(parsed) === 1) {
      features[0] = 100000; 
      features[3] = 1; 
      features[4] = 0; 
      features[14] = 50
    }

    return features;
  } catch (error) {
    console.error("Error extracting features:", error);
    // Return safe default (legitimate-looking features)
    return new Array(15).fill(0);
  }
}



/**
 * Feature 3: Shortining_Service
 * -1: URL uses a shortening service
 *  1: URL does not use a shortening service
 */
function Shortining_Service(url) {
  const hostname = url.hostname.toLowerCase();

  const shorteners = [
    "bit.ly",
    "goo.gl",
    "shorte.st",
    "go2l.ink",
    "x.co",
    "ow.ly",
    "tinyurl.com",
    "t.co",
    "tr.im",
    "is.gd",
    "cli.gs",
    "pic.gd",
    "DwarfURL.com",
    "yfrog.com",
    "migre.me",
    "ff.im",
    "tiny.cc",
    "url4.eu",
    "twit.ac",
    "su.pr",
    "twurl.nl",
    "snipurl.com",
    "BudURL.com",
    "short.to",
    "ping.fm",
    "post.ly",
    "Just.as",
    "bkite.com",
    "snipr.com",
    "fic.kr",
    "loopt.us",
    "doiop.com",
    "twitthis.com",
    "htxt.it",
    "AltURL.com",
    "RedirX.com",
    "DigBig.com",
    "short.ie",
    "u.mavrev.com",
    "kl.am",
    "wp.me",
    "rubyurl.com",
    "om.ly",
    "to.ly",
    "bit.do",
    "lnkd.in",
    "db.tt",
    "qr.ae",
    "adf.ly",
    "bitly.com",
    "cur.lv",
    "ity.im",
    "q.gs",
    "po.st",
    "bc.vc",
    "twitthis.com",
    "u.to",
    "j.mp",
    "buzurl.com",
    "cutt.us",
    "u.bb",
    "yourls.org",
    "prettylinkpro.com",
    "scrnch.me",
    "filoops.info",
    "vzturl.com",
    "qr.net",
    "1url.com",
    "tweez.me",
    "v.gd",
    "link.zip.net",
  ];

  for (const shortener of shorteners) {
    if (hostname === shortener || hostname.endsWith("." + shortener)) {
      return 1;
    }
  }

  return 0;
}

/**
 * Feature 4: having_At_Symbol
 * -1: URL contains @ symbol
 *  1: URL does not contain @ symbol
 */
function having_At_Symbol(url) {
  if (url.href.includes("@")) {
    return -1;
  }
  return 1;
}

/**
 * Feature 5: double_slash_redirecting
 * -1: The position of "//" in URL is > 7 (excluding protocol)
 *  1: The position of "//" in URL is <= 7
 */
function double_slash_redirecting(urlString) {
  // Find position of // after the protocol
  const protocolEnd = urlString.indexOf("://");
  if (protocolEnd === -1) {
    return 1;
  }

  const afterProtocol = urlString.substring(protocolEnd + 3);
  const doubleSlashPos = afterProtocol.indexOf("//");

  if (doubleSlashPos > 7) {
    return -1;
  }

  return 1;
}

/**
 * Feature 6: Prefix_Suffix
 * -1: Domain contains hyphen/dash
 *  1: Domain does not contain hyphen/dash
 */
function Prefix_Suffix(url) {
  const hostname = url.hostname;

  // Extract domain (remove subdomains)
  const parts = hostname.split(".");
  if (parts.length >= 2) {
    const domain = parts[parts.length - 2];
    if (domain.includes("-")) {
      return -1;
    }
  }

  return 1;
}

/**
 * Feature 7: having_Sub_Domain
 * -1: Having multiple sub-domains (>= 3 dots)
 *  0: Having one sub-domain (2 dots)
 *  1: No sub-domain (1 dot) or www only
 */
function having_Sub_Domain(url) {
  const hostname = url.hostname;
  const dots = (hostname.match(/\./g) || []).length;

  // Remove 'www' if present for counting
  let hostnameWithoutWww = hostname;
  if (hostname.startsWith("www.")) {
    hostnameWithoutWww = hostname.substring(4);
  }

  const dotsWithoutWww = (hostnameWithoutWww.match(/\./g) || []).length;

  if (dots >= 3) {
    return -1;
  } else if (dotsWithoutWww === 2 || dots === 2) {
    return 0;
  } else {
    return 1;
  }
}

/**
 * Feature 8: HTTPS_token
 * -1: HTTPS token in domain part (suspicious)
 *  1: No HTTPS token in domain part
 */
function HTTPS_token(url) {
  const hostname = url.hostname.toLowerCase();

  // Check if 'https' or 'http' appears in the domain name itself
  if (hostname.includes("https") || hostname.includes("http")) {
    return -1;
  }

  return 1;
}

/**
 * Feature 9: port
 * -1: Using non-standard port
 *  1: Using standard port (80, 443) or default
 */
function port(url) {
  const portNum = url.port;

  // Empty port means default (80 for http, 443 for https)
  if (portNum === "" || portNum === "80" || portNum === "443") {
    return 1;
  }

  return -1;
}

/**
 * Get feature names for display
 * @returns {Array<string>} Array of feature names
 */
function getFeatureNames() {
  return [
    "having_IP_Address",
    "URL_Length",
    "Shortining_Service",
    "having_At_Symbol",
    "double_slash_redirecting",
    "Prefix_Suffix",
    "having_Sub_Domain",
    "HTTPS_token",
    "port",
  ];
}

/**
 * Get human-readable feature explanations
 * @returns {Object} Map of feature names to explanations
 */
function getFeatureExplanations() {
  return {
    having_IP_Address: "Uses IP address instead of domain",
    URL_Length: "URL length analysis",
    Shortining_Service: "Uses URL shortening service",
    having_At_Symbol: "Contains @ symbol in URL",
    double_slash_redirecting: "Has suspicious // redirects",
    Prefix_Suffix: "Domain contains hyphens",
    having_Sub_Domain: "Number of subdomains",
    HTTPS_token: "HTTPS in domain name (suspicious)",
    port: "Uses non-standard port",
  };
}

// Export for use in other scripts
if (typeof module !== "undefined" && module.exports) {
  module.exports = {
    extractURLFeatures,
    getFeatureNames,
    getFeatureExplanations,
  };
}
