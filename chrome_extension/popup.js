/**
 * Popup UI Controller
 * Displays phishing detection results with beautiful interface
 */

// Feature names and explanations
const FEATURE_NAMES = [
  'having_IP_Address',
  'URL_Length',
  'Shortining_Service',
  'having_At_Symbol',
  'double_slash_redirecting',
  'Prefix_Suffix',
  'having_Sub_Domain',
  'HTTPS_token',
  'port'
];

const FEATURE_EXPLANATIONS = {
  'having_IP_Address': 'IP in URL',
  'URL_Length': 'URL Length',
  'Shortining_Service': 'URL Shortener',
  'having_At_Symbol': '@ Symbol',
  'double_slash_redirecting': '// Redirect',
  'Prefix_Suffix': 'Dash in Domain',
  'having_Sub_Domain': 'Subdomains',
  'HTTPS_token': 'HTTPS in Domain',
  'port': 'Non-standard Port'
};

/**
 * Get current tab URL
 */
async function getCurrentTab() {
  const [tab] = await chrome.tabs.query({ active: true, currentWindow: true });
  return tab;
}

/**
 * Display results in popup
 */
function displayResults(url, result) {
  const content = document.getElementById('content');
  
  // Determine status class
  let statusClass = 'status-unknown';
  let statusIcon = '❓';
  let statusText = 'Unknown';
  
  if (result.hybridProb !== undefined) {
    if (result.isPhishing) {
      statusClass = 'status-phishing';
      statusIcon = '🚨';
      statusText = 'PHISHING DETECTED';
    } else {
      statusClass = 'status-safe';
      statusIcon = '✅';
      statusText = 'SAFE';
    }
  }
  
  // Build probability bar classes
  const rfBarClass = result.rfProb > 0.9 ? 'prob-danger' : 
                     result.rfProb > 0.5 ? 'prob-warning' : 'prob-safe';
  const xgbBarClass = result.xgbProb > 0.9 ? 'prob-danger' : 
                      result.xgbProb > 0.5 ? 'prob-warning' : 'prob-safe';
  const hybridBarClass = result.hybridProb > 0.9 ? 'prob-danger' : 
                         result.hybridProb > 0.5 ? 'prob-warning' : 'prob-safe';
  
  // Build feature list
  let featuresHTML = '';
  if (result.features && result.features.length === 9) {
    featuresHTML = `
      <div class="features-section">
        <div class="features-title">🔍 Feature Analysis</div>
        ${result.features.map((value, idx) => {
          const name = FEATURE_NAMES[idx];
          const label = FEATURE_EXPLANATIONS[name];
          let valueClass = 'feature-safe';
          let valueText = 'OK';
          
          if (value === -1) {
            valueClass = 'feature-suspicious';
            valueText = 'SUSPICIOUS';
          } else if (value === 0) {
            valueClass = 'feature-neutral';
            valueText = 'NEUTRAL';
          }
          
          return `
            <div class="feature-item">
              <span class="feature-name">${label}</span>
              <span class="feature-value ${valueClass}">${valueText}</span>
            </div>
          `;
        }).join('')}
      </div>
    `;
  }
  
  content.innerHTML = `
    <div class="status-card ${statusClass}">
      <div class="status-icon">${statusIcon}</div>
      <div class="status-text">${statusText}</div>
    </div>
    
    <div class="url-display">
      <strong>URL:</strong> ${escapeHtml(url)}
    </div>
    
    <div class="metrics">
      <div class="metric-row">
        <span class="metric-label">🌲 Random Forest</span>
        <span class="metric-value">${(result.rfProb * 100).toFixed(2)}%</span>
      </div>
      <div class="prob-bar">
        <div class="prob-fill ${rfBarClass}" style="width: ${result.rfProb * 100}%"></div>
      </div>
      
      <div class="metric-row" style="margin-top: 12px;">
        <span class="metric-label">⚡ XGBoost</span>
        <span class="metric-value">${(result.xgbProb * 100).toFixed(2)}%</span>
      </div>
      <div class="prob-bar">
        <div class="prob-fill ${xgbBarClass}" style="width: ${result.xgbProb * 100}%"></div>
      </div>
      
      <div class="metric-row" style="margin-top: 12px;">
        <span class="metric-label">🤝 Hybrid Score</span>
        <span class="metric-value">${(result.hybridProb * 100).toFixed(2)}%</span>
      </div>
      <div class="prob-bar">
        <div class="prob-fill ${hybridBarClass}" style="width: ${result.hybridProb * 100}%"></div>
      </div>
    </div>
    
    ${featuresHTML}
    
    <button class="analyze-btn" id="reanalyze">
      🔄 Re-analyze Current Page
    </button>
    
    <div class="footer">
      Threshold: 90% | Models: RF + XGB | ONNX Runtime
    </div>
  `;
  
  // Add reanalyze button listener
  document.getElementById('reanalyze').addEventListener('click', async () => {
    const tab = await getCurrentTab();
    analyzeURL(tab.url);
  });
}

/**
 * Escape HTML to prevent XSS
 */
function escapeHtml(text) {
  const div = document.createElement('div');
  div.textContent = text;
  return div.innerHTML;
}

/**
 * Analyze a URL
 */
async function analyzeURL(url) {
  const content = document.getElementById('content');
  
  // Show loading
  content.innerHTML = `
    <div class="loading">
      <div class="spinner"></div>
      <p>Analyzing URL...</p>
    </div>
  `;
  
  try {
    // Send message to background script
    const result = await chrome.runtime.sendMessage({
      action: 'analyzeUrl',
      url: url
    });
    
    displayResults(url, result);
  } catch (error) {
    console.error('Error analyzing URL:', error);
    content.innerHTML = `
      <div class="status-card status-unknown">
        <div class="status-icon">⚠️</div>
        <div class="status-text">Error</div>
      </div>
      <div class="url-display">
        Failed to analyze URL. Please try again.
      </div>
      <button class="analyze-btn" id="retry">
        🔄 Retry
      </button>
    `;
    
    document.getElementById('retry').addEventListener('click', () => {
      analyzeURL(url);
    });
  }
}

/**
 * Initialize popup
 */
async function init() {
  try {
    // Get current tab
    const tab = await getCurrentTab();
    const url = tab.url;
    
    // Skip chrome:// URLs
    if (url.startsWith('chrome://') || 
        url.startsWith('chrome-extension://') ||
        url.startsWith('about:')) {
      const content = document.getElementById('content');
      content.innerHTML = `
        <div class="status-card status-unknown">
          <div class="status-icon">ℹ️</div>
          <div class="status-text">System Page</div>
        </div>
        <div class="url-display">
          This extension does not analyze Chrome system pages.
        </div>
      `;
      return;
    }
    
    // Try to get cached result first
    const cached = await chrome.storage.local.get(['lastUrl', 'lastResult', 'timestamp']);
    
    if (cached.lastUrl === url && 
        cached.lastResult && 
        (Date.now() - cached.timestamp) < 60000) {
      // Use cached result if less than 1 minute old
      displayResults(url, cached.lastResult);
    } else {
      // Analyze current URL
      await analyzeURL(url);
    }
  } catch (error) {
    console.error('Error initializing popup:', error);
  }
}

// Initialize when popup opens
document.addEventListener('DOMContentLoaded', init);
