package main

import (
	"context"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"unicode"
)

const ListenAddr = ":8026"

// ============================================================
// BLOCK PAGE
// ============================================================

const blockedHTML = `
<!DOCTYPE html>
<html>
<head>
<title>PhishShield</title>

<style>

body{
	background:#0f172a;
	color:white;
	font-family:Arial,sans-serif;
	display:flex;
	justify-content:center;
	align-items:center;
	height:100vh;
	margin:0;
}

.card{
	background:#1e293b;
	padding:40px;
	border-radius:18px;
	width:550px;
	text-align:center;
	box-shadow:0 0 25px rgba(0,0,0,0.5);
}

h1{
	color:#ef4444;
}

.url{
	color:#38bdf8;
	word-break:break-all;
	margin-top:20px;
	font-size:14px;
}

</style>
</head>

<body>

<div class="card">

<h1>⚠ Website Blocked</h1>

<p>
PhishShield detected this website as phishing.
</p>

<div class="url">
{{.URL}}
</div>

</div>

</body>
</html>
`

// ============================================================
// MAIN
// ============================================================

func main() {

	// =========================================================
	// LOG FILE SETUP
	// =========================================================

	os.MkdirAll("logs", os.ModePerm)

	logFile, err := os.OpenFile(
		"logs/phishshield.log",
		os.O_CREATE|
			os.O_WRONLY|
			os.O_APPEND,
		0666,
	)

	if err != nil {
		panic(err)
	}

	// Write logs to:
	// - terminal
	// - log file

	multiWriter := io.MultiWriter(
		os.Stdout,
		logFile,
	)

	log.SetOutput(multiWriter)

	log.SetFlags(
		log.Ldate |
			log.Ltime |
			log.Lshortfile,
	)

	log.Println("====================================")
	log.Println("Starting PhishShield Proxy")
	log.Println("====================================")

	// =========================================================
	// START SERVER
	// =========================================================

	server := &http.Server{
		Addr:    ListenAddr,
		Handler: http.HandlerFunc(proxyHandler),
	}

	go func() {

		log.Printf(
			"Proxy listening on %s\n",
			ListenAddr,
		)

		err := server.ListenAndServe()

		if err != nil &&
			err != http.ErrServerClosed {

			log.Fatalf(
				"Server Error: %v",
				err,
			)
		}
	}()

	waitForShutdown(server)
}

// ============================================================
// PROXY HANDLER
// ============================================================

func proxyHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	log.Printf(
		"[%s] %s\n",
		r.Method,
		r.Host,
	)

	// HTTPS CONNECT
	if r.Method == http.MethodConnect {

		handleHTTPS(w, r)

		return
	}

	targetURL := r.URL.String()

	if targetURL == "" {

		targetURL =
			"http://" +
				r.Host +
				r.URL.Path
	}

	isPhishing := predictURL(targetURL)

	if isPhishing {

		log.Printf(
			"[BLOCKED] %s\n",
			targetURL,
		)

		showBlockedPage(w, targetURL)

		return
	}

	log.Printf(
		"[ALLOWED] %s\n",
		targetURL,
	)

	forwardRequest(w, r)
}

// ============================================================
// HTTPS CONNECT
// ============================================================

func handleHTTPS(
	w http.ResponseWriter,
	r *http.Request,
) {

	target := r.Host

	log.Printf(
		"[HTTPS CONNECT] %s\n",
		target,
	)

	isPhishing := predictURL(
		"https://" + target,
	)

	if isPhishing {

		log.Printf(
			"[BLOCKED HTTPS] %s\n",
			target,
		)

		showBlockedPage(w, target)

		return
	}

	destConn, err := net.DialTimeout(
		"tcp",
		target,
		10*time.Second,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusServiceUnavailable,
		)

		return
	}

	hijacker, ok := w.(http.Hijacker)

	if !ok {

		http.Error(
			w,
			"Hijacking unsupported",
			http.StatusInternalServerError,
		)

		return
	}

	clientConn, _, err := hijacker.Hijack()

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusServiceUnavailable,
		)

		return
	}

	_, err = clientConn.Write(
		[]byte(
			"HTTP/1.1 200 Connection Established\r\n\r\n",
		),
	)

	if err != nil {

		clientConn.Close()
		destConn.Close()

		return
	}

	go tunnel(destConn, clientConn)
	go tunnel(clientConn, destConn)
}

func tunnel(
	dst io.WriteCloser,
	src io.ReadCloser,
) {

	defer dst.Close()
	defer src.Close()

	io.Copy(dst, src)
}

// ============================================================
// FORWARD REQUEST
// ============================================================

func forwardRequest(
	w http.ResponseWriter,
	r *http.Request,
) {

	req, err := http.NewRequest(
		r.Method,
		r.URL.String(),
		r.Body,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	req.Header = r.Header

	resp, err := http.DefaultTransport.RoundTrip(req)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusServiceUnavailable,
		)

		return
	}

	defer resp.Body.Close()

	copyHeaders(w.Header(), resp.Header)

	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)
}

func copyHeaders(
	dst,
	src http.Header,
) {

	for k, vv := range src {

		for _, v := range vv {

			dst.Add(k, v)
		}
	}
}

// ============================================================
// BLOCK PAGE
// ============================================================

func showBlockedPage(
	w http.ResponseWriter,
	blockedURL string,
) {

	tmpl := template.Must(
		template.New("blocked").Parse(blockedHTML),
	)

	w.Header().Set(
		"Content-Type",
		"text/html",
	)

	w.WriteHeader(http.StatusForbidden)

	tmpl.Execute(
		w,
		map[string]string{
			"URL": blockedURL,
		},
	)
}

// ============================================================
// ENSEMBLE PREDICTION
// ============================================================

func predictURL(rawURL string) bool {

	features32 := ExtractUrlFeatures(rawURL)

	log.Printf(
		"Features: %+v\n",
		features32,
	)

	// float32 -> float64
	features := make([]float64, len(features32))

	for i, v := range features32 {

		features[i] = float64(v)
	}

	// ====================================================
	// RANDOM FOREST PREDICTION
	// ====================================================

	rfResult := PredictRF(features)

	// ====================================================
	// XGBOOST PREDICTION
	// ====================================================
	xgbResult := PredictXGB(features)

	// Use phishing probability
	rfProb := rfResult[1]
	xgbProb := xgbResult[1]

	finalProb :=
		(rfProb + (xgbProb * 1.2)) / 2.2

	log.Printf(
		"RF=%f XGB=%f FINAL=%f URL=%s\n",
		rfProb,
		xgbProb,
		finalProb,
		rawURL,
	)

	return finalProb < 0.07 // **** ^ maybe bump up?
}

// ============================================================
// FEATURE EXTRACTION
// MATCHES TRAINING FEATURES EXACTLY
// ============================================================

// ============================================================
// ADVANCED URL FEATURE EXTRACTION
// ============================================================

func ExtractUrlFeatures(rawURL string) []float32 {

	// Safe default
	defaultFeatures := make([]float32, 15)

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Feature extraction panic: %v", r)
		}
	}()

	rawURL = strings.TrimSpace(rawURL)

	if rawURL == "" {
		return defaultFeatures
	}

	// Ensure protocol exists
	if !strings.HasPrefix(
		strings.ToLower(rawURL),
		"http://",
	) &&
		!strings.HasPrefix(
			strings.ToLower(rawURL),
			"https://",
		) {

		rawURL = "http://" + rawURL
	}

	parsed, err := url.Parse(rawURL)

	if err != nil {
		log.Printf("URL Parse Error: %v", err)
		return defaultFeatures
	}

	domain := strings.ToLower(parsed.Hostname())
	href := strings.ToLower(parsed.String())
	pathname := strings.ToLower(parsed.Path)
	search := strings.ToLower(parsed.RawQuery)

	// ========================================================
	// BASIC FEATURES
	// ========================================================

	urlLength := float32(len(href))

	domainLength := float32(len(domain))

	noOfSubDomain := float32(
		strings.Count(domain, ".") - 1,
	)

	if noOfSubDomain < 0 {
		noOfSubDomain = 0
	}

	isDomainIP := float32(0)

	if isIPAddress(domain) {
		isDomainIP = 1
	}

	isHTTPS := float32(0)

	if strings.HasPrefix(href, "https://") {
		isHTTPS = 1
	}

	// ========================================================
	// COUNTERS
	// ========================================================

	var noOfDigits float32
	var noOfLetters float32
	var noOfOtherSpecial float32

	for _, c := range href {

		if unicode.IsDigit(c) {

			noOfDigits++

		} else if unicode.IsLetter(c) {

			noOfLetters++

		} else if !strings.ContainsAny(
			string(c),
			".:/?=&-_",
		) {

			noOfOtherSpecial++
		}
	}

	noOfEquals := float32(
		strings.Count(href, "="),
	)

	noOfQMark := float32(
		strings.Count(href, "?"),
	)

	noOfAmpersand := float32(
		strings.Count(href, "&"),
	)

	// ========================================================
	// RATIOS
	// ========================================================

	var specialRatio float32
	var letterRatio float32
	var digitRatio float32

	if urlLength > 0 {

		specialRatio =
			noOfOtherSpecial / urlLength

		letterRatio =
			noOfLetters / urlLength

		digitRatio =
			noOfDigits / urlLength
	}

	// ========================================================
	// TLD
	// ========================================================

	tld := ""

	parts := strings.Split(domain, ".")

	if len(parts) > 0 {
		tld = parts[len(parts)-1]
	}

	tldLength := float32(len(tld))

	// ========================================================
	// FEATURE VECTOR
	// ========================================================

	features := []float32{
		urlLength,        // 0
		domainLength,     // 1
		noOfSubDomain,    // 2
		isDomainIP,       // 3
		isHTTPS,          // 4
		noOfDigits,       // 5
		noOfOtherSpecial, // 6
		noOfEquals,       // 7
		noOfQMark,        // 8
		noOfAmpersand,    // 9
		specialRatio,     // 10
		letterRatio,      // 11
		digitRatio,       // 12
		noOfLetters,      // 13
		tldLength,        // 14
	}

	// ========================================================
	// SUSPICIOUS TLD CHECK
	// ========================================================

	suspiciousTLDs := map[string]bool{
		"top":     true,
		"xyz":     true,
		"online":  true,
		"site":    true,
		"club":    true,
		"shop":    true,
		"store":   true,
		"live":    true,
		"digital": true,
		"buzz":    true,
		"fun":     true,
		"monster": true,
		"loan":    true,
		"win":     true,
		"men":     true,
		"gq":      true,
		"cf":      true,
		"ml":      true,
		"ga":      true,
		"tk":      true,
	}

	tldSuspicious := suspiciousTLDs[tld]

	// ========================================================
	// SCAM KEYWORDS
	// ========================================================

	scamKeywords := []string{
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
	}

	keywordMatch := false

	for _, kw := range scamKeywords {

		if strings.Contains(domain, kw) ||
			strings.Contains(pathname, kw) ||
			strings.Contains(search, kw) {

			keywordMatch = true
			break
		}
	}

	// ========================================================
	// PUNYCODE
	// ========================================================

	isPuny := isPunycodeDomain(domain)

	// ========================================================
	// HOMOGRAPH
	// ========================================================

	isHomograph := hasSuspiciousHomograph(domain)

	// ========================================================
	// SHORTENER
	// ========================================================

	isShortened := ShortiningService(domain)

	// ========================================================
	// BOOST PHISHING SIGNALS
	// ========================================================

	if tldSuspicious ||
		keywordMatch ||
		isPuny ||
		isHomograph ||
		isShortened {

		features[0] = 100000
		features[3] = 1
		features[4] = 0
		features[14] = 50
	}

	return features
}

// ============================================================
// PUNYCODE CHECK
// ============================================================

func isPunycodeDomain(
	hostname string,
) bool {

	return strings.Contains(
		strings.ToLower(hostname),
		"xn--",
	)
}

// ============================================================
// HOMOGRAPH CHECK
// ============================================================

func hasSuspiciousHomograph(
	hostname string,
) bool {

	// Mixed script detection

	hasLatin := false
	hasCyrillic := false
	hasGreek := false

	for _, c := range hostname {

		switch {

		case unicode.In(c, unicode.Latin):
			hasLatin = true

		case unicode.In(c, unicode.Cyrillic):
			hasCyrillic = true

		case unicode.In(c, unicode.Greek):
			hasGreek = true
		}
	}

	// Mixed scripts are suspicious
	if hasLatin &&
		(hasCyrillic || hasGreek) {

		return true
	}

	// Common homograph chars
	suspiciousChars := []rune{
		'а',
		'ɑ',
		'с',
		'ϲ',
		'е',
		'і',
		'ο',
		'р',
		'х',
		'у',
		'γ',
		'ӏ',
		'ł',
	}

	for _, c := range hostname {

		for _, s := range suspiciousChars {

			if c == s {
				return true
			}
		}
	}

	return false
}

// ============================================================
// URL SHORTENER DETECTION
// ============================================================

func ShortiningService(
	hostname string,
) bool {

	shorteners := map[string]bool{

		"bit.ly":            true,
		"goo.gl":            true,
		"tinyurl.com":       true,
		"t.co":              true,
		"ow.ly":             true,
		"is.gd":             true,
		"buff.ly":           true,
		"adf.ly":            true,
		"bitly.com":         true,
		"cutt.ly":           true,
		"rebrand.ly":        true,
		"tiny.cc":           true,
		"shorturl.at":       true,
		"lnkd.in":           true,
		"db.tt":             true,
		"qr.ae":             true,
		"v.gd":              true,
		"1url.com":          true,
		"yourls.org":        true,
		"prettylinkpro.com": true,
	}

	hostname = strings.ToLower(hostname)

	if shorteners[hostname] {
		return true
	}

	for shortener := range shorteners {

		if strings.HasSuffix(
			hostname,
			"."+shortener,
		) {

			return true
		}
	}

	return false
}

// ============================================================
// IP ADDRESS CHECK
// ============================================================

func isIPAddress(s string) bool {

	return strings.Count(s, ".") == 3 &&
		!strings.ContainsAny(
			s,
			"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		)
}

// ============================================================
// CLEAN SHUTDOWN
// ============================================================

func waitForShutdown(
	server *http.Server,
) {

	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-stop

	log.Println("Shutting down proxy...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {

		log.Printf(
			"Shutdown Error: %v\n",
			err,
		)
	}

	log.Println("PhishShield stopped")
}
