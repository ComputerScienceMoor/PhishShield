# PhishShield Browser Extension — AI-Powered Phishing Detection

## Overview

**PhishShield Browser Extension** is a Machine Learning-powered phishing detection extension developed as part of the **PhishShield HND Final Year Project**.

The extension analyzes URLs in real time using lexical URL features and an AI phishing detection model exported to ONNX format.

When a suspicious or phishing website is detected, the extension blocks access and warns the user before the page loads.

---

# Features

* Real-time phishing detection
* Browser-based URL analysis
* ONNX Machine Learning inference
* Random Forest + XGBoost support
* Lightweight lexical feature extraction
* Suspicious keyword detection
* Punycode detection
* Homograph attack detection
* URL shortener detection
* Fast local inference
* No cloud dependency
* Offline functionality

---

# System Architecture

```text id="f93ob0"
Browser URL
     │
     ▼
Feature Extraction
     │
     ▼
ONNX Model Inference
     │
     ▼
Prediction
 ┌───────────────┐
 │ Legitimate    │ → Allow Access
 └───────────────┘

 ┌───────────────┐
 │ Phishing      │ → Block Page
 └───────────────┘
```

---

# Technologies Used

## Frontend

* JavaScript
* Chrome Extensions

## Machine Learning

* ONNX Runtime Web
* Scikit-learn
* XGBoost

## Model Conversion

* skl2onnx
* onnxmltools

---

# Dataset

The phishing dataset used for training:

[PhiUSIIL Phishing URL Dataset](https://github.com/elaaatif/DATA-MINING-PhiUSIIL-Phishing-URL)

---

# Project Structure

```text id="h7es09"
extension/
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

# How The Extension Works

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

# URL Features Used

The extension extracts 15 lexical features from every URL.

| Feature                    | Description              |
| -------------------------- | ------------------------ |
| URLLength                  | Total URL length         |
| DomainLength               | Domain length            |
| NoOfSubDomain              | Number of subdomains     |
| IsDomainIP                 | Whether domain is an IP  |
| IsHTTPS                    | Whether HTTPS is used    |
| NoOfDegitsInURL            | Number of digits         |
| NoOfOtherSpecialCharsInURL | Special characters count |
| NoOfEqualsInURL            | "=" count                |
| NoOfQMarkInURL             | "?" count                |
| NoOfAmpersandInURL         | "&" count                |
| SpacialCharRatioInURL      | Special character ratio  |
| LetterRatioInURL           | Letter ratio             |
| DegitRatioInURL            | Digit ratio              |
| NoOfLettersInURL           | Total letters            |
| TLDLength                  | Top-level domain length  |

---

# Additional Security Checks

The extension also performs additional lightweight phishing heuristics:

* Suspicious TLD detection
* Scam keyword detection
* Punycode detection
* Homograph attack detection
* URL shortener detection

These heuristics strengthen phishing detection accuracy.

---

# Machine Learning Models

The extension uses:

* Random Forest
* XGBoost
* Ensemble Voting

The models were trained in Python and exported to ONNX format for browser inference.

---

# Step 1 — Install Python Dependencies

Install required ML libraries:

```bash id="i48a1j"
pip install pandas numpy scikit-learn xgboost skl2onnx onnxmltools jupyter
```

---

# Step 2 — Run The Training Notebook

Start Jupyter Notebook:

```bash id="06hys8"
jupyter notebook
```

Open:

```text id="omx1d7"
training/extension_ml_training_notebook.ipynb
```

Run all cells sequentially.

The notebook will:

* Load dataset
* Train ML models
* Evaluate accuracy
* Export ONNX models

Generated files:

```text id="2hrq5h"
phishshield_rf.onnx
phishshield_xgb.onnx
```

Move these files into:

```text id="g8k4gb"
models/
```

---

# Step 3 — Load Extension Into Chrome

## Open Chrome Extensions Page

Open:

```text id="6qs9hq"
chrome://extensions
```

OR Open:

```text id="6qs9hq"
brave://extensions/
```

---

## Enable Developer Mode

Turn ON:

```text id="9tjlwm"
Developer Mode
```

(top-right corner)

---

## Load The Extension

Click:

```text id="smmh7y"
Load unpacked
```

Select the extension project folder:

```text id="4u0nrf"
extension/
```

The extension should now appear in Chrome.

---

# Step 4 — Test The Extension

## Legitimate Websites

Visit:

* [https://google.com](https://google.com)
* [https://github.com](https://github.com)
* [https://wikipedia.org](https://wikipedia.org)

They should load normally.

---

## Test Suspicious URLs

Try URLs such as:

```text id="xot8u6"
http://paypal-login-secure-account-update.com
```

Expected behavior:

* Warning page displayed
* Navigation blocked

---

# Example Detection Flow

```text id="d8utlc"
Visited URL:
http://secure-paypal-login-verification.xyz

↓
Feature Extraction

↓
ONNX Inference

↓
Prediction:
Phishing Probability = High!

↓
Blocked
```

---

# Browser Permissions

The extension requires permissions such as:

```json id="ru76jf"
{
  "permissions": [
    "tabs",
    "webNavigation",
    "storage"
  ]
}
```

These permissions allow the extension to:

* monitor URLs
* analyze navigation
* block malicious pages

---

# Advantages Of Browser-Based Detection

* Real-time protection
* Fast local inference
* Offline capability
* No cloud latency
* No external API dependency
* Lightweight execution

---

# Limitations

* Lexical analysis only
* Does not inspect webpage content
* No DNS reputation analysis
* No TLS certificate analysis
* False positives may occur

---

# Future Improvements

Potential future enhancements:

* Deep Learning models
* Transformer-based phishing detection
* Threat intelligence integration
* Cloud synchronization
* Browser sync support
* Real-time blacklist updates
* User reporting dashboard

---

# Research Contribution

This project demonstrates how browser extensions combined with Machine Learning and ONNX inference can provide lightweight real-time phishing protection directly within the browser environment.

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

This project is intended for educational and research purposes only.
