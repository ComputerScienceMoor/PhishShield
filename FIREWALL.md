# PhishShield — AI-Powered Phishing Detection Firewall/Proxy

## Overview

**PhishShield** is a lightweight AI-powered phishing detection system developed as a Higher National Diploma (HND) final year project.

The system works as a local HTTP/HTTPS proxy server that intercepts web traffic in real time, extracts lexical URL features, and uses a Machine Learning ensemble model (**Random Forest + XGBoost**) to determine whether a URL is phishing or legitimate.

If a phishing URL is detected, the request is blocked.

---

# Features

* Real-time phishing URL detection
* HTTP proxy support
* HTTPS CONNECT tunneling support
* Ensemble Machine Learning detection
* Go backend
* Lightweight and fast
* Detailed request logging
* URL lexical feature extraction
* Suspicious keyword detection
* Punycode & homograph attack detection
* URL shortener detection
* Cross-platform support

---

# Project Architecture

```text
Browser
   │
   ▼
PhishShield Go Firewall/Proxy
   │
   ├── URL Feature Extraction
   │
   ├── Random Forest Model
   │
   ├── XGBoost Model
   │
   └── Ensemble Voting
   │
   ▼
Decision
   ├── Legitimate → Forward Request
   └── Phishing → Block Page
```

---

# Technologies Used

## Backend

* Go

## Machine Learning

* Scikit-learn
* XGBoost
* m2cgen

## Data Science

* Jupyter Notebook
* Pandas
* NumPy

---

# Dataset

The phishing dataset used for training was obtained from:

[PhiUSIIL Phishing URL Dataset](https://github.com/elaaatif/DATA-MINING-PhiUSIIL-Phishing-URL)

---

# How The System Works

1. User opens a website in the browser
2. Browser sends traffic through the local proxy
3. PhishShield intercepts the request
4. URL lexical features are extracted
5. Features are passed into:

   * Random Forest model
   * XGBoost model
6. Predictions are combined using weighted ensemble voting
7. If phishing probability exceeds threshold:

   * Website is blocked
8. Otherwise:

   * Request is forwarded normally

---

# URL Features Used

The system extracts 15 lexical features from every URL:

| Feature                    | Description              |
| -------------------------- | ------------------------ |
| URLLength                  | Total URL length         |
| DomainLength               | Domain name length       |
| NoOfSubDomain              | Number of subdomains     |
| IsDomainIP                 | Whether domain is an IP  |
| IsHTTPS                    | Whether HTTPS is used    |
| NoOfDegitsInURL            | Number of digits         |
| NoOfOtherSpecialCharsInURL | Special characters count |
| NoOfEqualsInURL            | Number of "="            |
| NoOfQMarkInURL             | Number of "?"            |
| NoOfAmpersandInURL         | Number of "&"            |
| SpacialCharRatioInURL      | Ratio of special chars   |
| LetterRatioInURL           | Ratio of letters         |
| DegitRatioInURL            | Ratio of digits          |
| NoOfLettersInURL           | Total letters            |
| TLDLength                  | Top-level domain length  |

---

# Additional Security Checks

The system also includes lightweight phishing heuristics:

* Suspicious TLD detection
* Scam keyword detection
* Punycode detection
* Homograph attack detection
* URL shortener detection

These checks boost phishing probability without changing model dimensions.

---

# Project Structure

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

# Step 1 — Install Requirements

## Install Go

Download and install Go:

[Go Programming Language](https://go.dev/dl/)

Verify installation:

```bash
go version
```

---

## Install Python

Download Python:

[Python Downloads](https://www.python.org/downloads/)

---

## Install Jupyter Notebook

```bash
pip install notebook
```

---

# Step 2 — Install Python Dependencies

Install all ML dependencies:

```bash
pip install pandas numpy scikit-learn xgboost matplotlib seaborn joblib m2cgen
```

---

# Step 3 — Run The Training Notebook

Start Jupyter:

```bash
jupyter notebook
```

Open:

```text
training/firewall_ml_training_notebook.ipynb
```

Run all notebook cells sequentially.

The notebook will:

* Load dataset
* Train Random Forest model
* Train XGBoost model
* Build ensemble
* Evaluate accuracy
* Export models to pure Go code

---

# Step 4 — Generate Go ML Models

The notebook automatically generates:

```text
rf_model.go
xgb_model.go
```

---

# Step 5 — Build The Go Proxy

Inside the "firewall" project folder directory on the terminal/CLI, run:

```bash
go mod tidy
go run .
```

OR

```bash
go mod tidy
go build
```

This generates:

```text
phishshield
```

---

# Step 6 — Run The Proxy

Start the proxy server:

```bash
./phishshield
```

Expected output:

```text
Starting PhishShield Proxy
Proxy listening on :8026
```

---

# Step 7 — Configure Browser/System Proxy

## macOS

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

## Windows

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

# Step 8 — Test The System

## Legitimate Websites

Visit:

* [https://google.com](https://google.com)
* [https://github.com](https://github.com)
* [https://wikipedia.org](https://wikipedia.org)

They should load normally.

---

## Phishing Test URLs

Try suspicious URLs such as:

```text
http://paypal-login-secure-account-update.com
```

Expected behavior:

* Request blocked
* Warning page displayed
* Logs written

---

# Logging

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

# How Ensemble Voting Works

The system combines both model predictions:

```text
Final Probability =
(RandomForest + XGBoost × 1.2) / 2.2
```

XGBoost is slightly weighted higher because it performed better during evaluation.

---

# Advantages Of The System

* Lightweight deployment
* No external APIs
* No cloud dependency
* Real-time protection
* Explainable lexical analysis
* Cross-platform
* Easy to maintain

---

# Limitations

* Lexical-only detection
* Does not inspect webpage DOM
* Does not analyze page content
* No DNS reputation analysis
* HTTPS traffic is tunneled, not decrypted

---

# Future Improvements

Possible future enhancements include:

* Deep Learning integration
* Browser extension integration
* Threat intelligence feeds
* Real-time blacklist synchronization
* Cloud dashboard
* TLS certificate analysis
* DNS reputation scoring

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
