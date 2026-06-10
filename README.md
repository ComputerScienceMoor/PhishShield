# PhishShield - FARAYOLA, TIMILEYIN EMMANUEL - HNDNCC/24/013

**PhishShield** is a lightweight AI-powered phishing detection system developed as a Higher National Diploma (HND) final year project. It uses a Machine Learning ensemble model (**Random Forest + XGBoost**) to determine whether a URL is phishing or legitimate. If a phishing URL is detected, the request is blocked.

A Machine Learning (ML) approach to preventing Phising. (My Federal College Of Animal Health & Production Technology, Moor-Plantation, Ibadan (FCAHPT). Higher National Diploma (HND) Computer Science [Networking and Cloud Computing (NCC)] 2026 Final Year Project)




## Whole project folder structure
```text id="h7es09"
phishshield/
│
├── datasets/           => [Contains the PhiUSIL URL dataset used to train the ML models]
├── chrome_extension/   => [Contains the chrome extension implementation code]
├── firewall/           => [Contains the firewall implementation code]
├── training/           => [Contains the Jupyter notebooks used to code the ML models]
├── .gitignore          => [to ignore pushing unwanted files to the repo]
├── EXTENSION.md        => [Chrome extension implementation documentation]
├── FIREWALL.md         => [Firewall / Proxy implementation documentation]
├── LICENSE             => [License]
├── README.md           => [Project implementation documentation]
```


## 1) PhishShield Browser Extension - Usage
### Project Structure

```text id="h7es09"
chrome_extension/
│
├── lib/
    ├── ort-wasm-simd.wasm
│   ├── ort.min.js
│   └── sirens.mp3
├── models/
│   ├── phishshield_rf.onnx
│   └── phishshield_xgb.onnx
├── background.js
├── blocked.html
├── blocked.js
├── features.js
├── manifest.json
├── popup.html
├── popup.js
```

---
### How The Extension Works

1. User visits a website
2. Extension intercepts the URL
3. Lexical URL features are extracted
4. Features are passed into the ONNX models
5. AI predicts whether the URL is phishing or legitimate
6. If phishing:

   * User is redirected to a warning page
7. Otherwise:

   * Browsing continues normally

---

### Load Extension Into Chrome

#### Launch Chrome / Brave Web Browser and Open Chrome Extensions Page

Open:

```text id="6qs9hq"
chrome://extensions
```

OR Open:

```text id="6qs9hq"
brave://extensions/
```

---

#### Enable Developer Mode

Turn ON:

```text id="9tjlwm"
Developer Mode
```

(top-right corner)

---

#### Load The Extension

Click:

```text id="smmh7y"
Load unpacked
```

Select the chrome_extension project folder:

```text id="4u0nrf"
chrome_extension/
```

The extension should now appear in Chrome.

---

### Test The Extension

#### Legitimate Websites

Visit:

* [https://google.com](https://google.com)
* [https://github.com](https://github.com)
* [https://wikipedia.org](https://wikipedia.org)

They should load normally.

---

#### Test Suspicious URLs

Try URLs such as:

```text id="xot8u6"
http://paypal-login-secure-account-update.com
```

Expected behavior:

* Warning page displayed
* Navigation blocked

---
Check [EXTENSION.md](EXTENSION.md) to read more documentation on the above.

---
---
---
---
---

## 2) PhishShield AI-Powered Phishing Detection Firewall / Proxy - Usage
---

### Project Structure

```text
firewall/
│
├── main.go
├── rf_model.go
├── xgb_model.go
├── go.mod
├── go.sum
├── logs/
│   └── phishshield.log
 
```

---

### Start / Run The Proxy

Start the proxy server:

Change directory (cd) into the "firewall" project folder directory on the terminal/CLI, then run the commands below:

```bash
go mod tidy
go run .
```

Expected output:

```text
Starting PhishShield Proxy
Proxy listening on :8026
```

---

### Configure Browser / System Proxy

#### On macOS (Apple machines)

Open:

```text
System Settings
→ Network
→ Wi-Fi
→ Details
→ Proxies
```

Enable:

* Web Proxy (HTTP)
* Secure Web Proxy (HTTPS)

Use:

```text
Server: 127.0.0.1
Port: 8026
```

Click:

* OK
* Apply

---

#### Windows machines

Open:

```text
Settings
→ Network & Internet
→ Proxy
```

Enable manual proxy setup:

```text
Address: 127.0.0.1
Port: 8026
```

---

### Test The System

#### Legitimate Websites

Visit:

* [https://google.com](https://google.com)
* [https://github.com](https://github.com)
* [https://wikipedia.org](https://wikipedia.org)

They should load normally.

---

#### Phishing Test URLs

Try suspicious URLs such as:

```text
http://paypal-login-secure-account-update.com
```

Expected behavior:

* Request blocked
* Warning page displayed
* Logs written

---

### Logging

All logs are stored in:

```text
logs/phishshield.log
```

Example log output:

```text
2026/05/28 12:44:01 [CONNECT] google.com:443
2026/05/28 12:44:01 RF=0.01 XGB=0.03 FINAL=0.02
2026/05/28 12:44:01 [ALLOWED] https://google.com
```

---

Check [FIREWALL.md](FIREWALL.md) to read more documentation on the above.

---
---
---
---
---


# Research Contribution

This project demonstrates that lightweight lexical feature extraction combined with ensemble machine learning can provide effective phishing detection without requiring heavy browser instrumentation or cloud-based analysis.

---

# Author

Developed by:

**Timileyin Farayola**
  * [https://linkedin.com/in/timileyin-farayola](https://linkedin.com/in/timileyin-farayola)
  * [https://github.com/rafmme](https://github.com/rafmme)
  * [https://x.com/rafmme](https://x.com/rafmme)


Federal College Of Animal Health & Production Technology, Moor-Plantation, Ibadan (FCAHPT). Higher National Diploma (HND) 2026 Final Year Project

---

# License

This project is for educational and research purposes.