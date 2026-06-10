/**
 * Background Service Worker for Phishing Detection
 * Handles navigation events and runs ONNX inference
 */

// Import feature extraction
importScripts("features.js");
importScripts("lib/ort.min.js");

// Global variables for ONNX sessions
let rfSession = null;
let xgbSession = null;
let isModelLoaded = false;

// Detection threshold
const PHISHING_THRESHOLD = 0.1;

/**
 * Load ONNX models on startup
 */
async function loadModels() {
  try {
    console.log("Loading ONNX models...");
    ort.env.wasm.proxy = false;
    ort.env.wasm.numThreads = 1;
    ort.env.wasm.wasmPaths = chrome.runtime.getURL("lib/");

    const options = { executionProviders: ["wasm"] };

    // Load Random Forest model
    const rfModelPath = chrome.runtime.getURL("./models/phishshield_rf.onnx");
    rfSession = await ort.InferenceSession.create(rfModelPath, options);
    console.log("✓ Random Forest model loaded");

    // Load XGBoost model
    const xgbModelPath = chrome.runtime.getURL("./models/phishshield_xgb.onnx");
    xgbSession = await ort.InferenceSession.create(xgbModelPath, options);
    console.log("✓ XGBoost model loaded");

    isModelLoaded = true;
    console.log("All models loaded successfully");
  } catch (error) {
    console.error("Error loading models:", error);
    isModelLoaded = false;
  }
}

/**
 * Run inference on both models and return hybrid prediction
 * @param {string} url - URL to analyze
 * @returns {Object} Prediction results
 */
 async function predictPhishing(url) {
  if (!isModelLoaded || !rfSession || !xgbSession) {
    console.warn("Models not loaded yet");
    return { isPhishing: false, rfProb: 0, xgbProb: 0, hybridProb: 0, features: [] };
  }

  try {
    const features = extractURLFeatures(url);
    console.log("Extracted features for", url, ":", features);

    const inputTensor = new ort.Tensor("float32", new Float32Array(features), [1, 15]);

    const rfInputName = rfSession.inputNames[0];
    const xgbInputName = xgbSession.inputNames[0];

    const rfOutputs = await rfSession.run({ [rfInputName]: inputTensor });
    const xgbOutputs = await xgbSession.run({ [xgbInputName]: inputTensor });

    // Debug output (keep for now)
    console.log("RF raw output:", rfOutputs);
    console.log("XGB raw output:", xgbOutputs);

    let rfLegitProb = 0.5;
    let xgbLegitProb = 0.5;

    // === RF Model - Safe extraction ===
    if (rfOutputs.probabilities && rfOutputs.probabilities.data) {
      rfLegitProb = rfOutputs.probabilities.data[1] || 0.5;   // index 1 = Legitimate
    } else if (rfOutputs[0] && rfOutputs[0].data) {
      rfLegitProb = rfOutputs[0].data[1] || 0.5;
    }

    // === XGB Model - Safe extraction ===
    if (xgbOutputs.probabilities && xgbOutputs.probabilities.data) {
      xgbLegitProb = xgbOutputs.probabilities.data[1] || 0.5;
    } else if (xgbOutputs[0] && xgbOutputs[0].data) {
      xgbLegitProb = xgbOutputs[0].data[1] || 0.5;
    }

    // Soft voting - probability of being Legitimate
    const ensembleLegitProb = (rfLegitProb + 1.2 * xgbLegitProb) / (1 + 1.2);

    // Adjusted threshold to reduce false positives on legitimate sites
    const isPhishing = ensembleLegitProb < PHISHING_THRESHOLD;

    console.log(`RF Legit: ${rfLegitProb.toFixed(4)} | XGB Legit: ${xgbLegitProb.toFixed(4)} | Ensemble Legit: ${ensembleLegitProb.toFixed(4)} → ${isPhishing ? 'PHISHING' : 'LEGITIMATE'}`);

    return {
      isPhishing,
      rfProb: rfLegitProb,
      xgbProb: xgbLegitProb,
      hybridProb: ensembleLegitProb,
      features,
    };

  } catch (error) {
    console.error("Error during prediction:", error);
    return { isPhishing: false, rfProb: 0, xgbProb: 0, hybridProb: 0, features: [] };
  }
}

/**
 * Handle navigation events
 */
chrome.webNavigation.onBeforeNavigate.addListener(async (details) => {
  // Only process main frame navigations
  if (details.frameId !== 0) {
    return;
  }

  const url = details.url;

  // Skip chrome:// and extension pages
  if (
    url.startsWith("chrome://") ||
    url.startsWith("chrome-extension://") ||
    url.startsWith("about:")
  ) {
    return;
  }

  console.log("Analyzing URL:", url);

  // Run phishing detection
  const result = await predictPhishing(url);

  // Store result for popup
  await chrome.storage.local.set({
    lastUrl: url,
    lastResult: result,
    timestamp: Date.now(),
  });

  // Block if phishing detected
  if (result.isPhishing) {
    console.warn("⚠ PHISHING DETECTED:", url);

    // Redirect to blocked page
    const blockedUrl =
      chrome.runtime.getURL("blocked.html") +
      "?url=" +
      encodeURIComponent(url) +
      "&prob=" +
      result.hybridProb.toFixed(4);

    chrome.tabs.update(details.tabId, { url: blockedUrl });
  } else {
    console.log("✓ URL appears legitimate");
  }
});

/**
 * Handle messages from popup
 */
chrome.runtime.onMessage.addListener((request, sender, sendResponse) => {
  if (request.action === "analyzeUrl") {
    predictPhishing(request.url).then((result) => {
      sendResponse(result);
    });
    return true; // Keep channel open for async response
  }

  if (request.action === "getModelStatus") {
    sendResponse({ isLoaded: isModelLoaded });
    return true;
  }
});

/**
 * Initialize models on extension install/update
 */
chrome.runtime.onInstalled.addListener(() => {
  console.log("Extension installed/updated");
  loadModels();
});

/**
 * Load models on startup
 */
loadModels();

console.log("AI Phishing Detector background service worker started");
