// Get URL parameters
const params = new URLSearchParams(window.location.search);
const blockedUrl = params.get("url") || "Unknown URL";
const probability = params.get("prob") || "0.0000";

// Display blocked URL
document.getElementById("blockedUrl").textContent =
  decodeURIComponent(blockedUrl);
document.getElementById("probability").textContent =
  (parseFloat(probability) * 100).toFixed(2) + "%";

// Go back button
document.getElementById("goBack").addEventListener("click", () => {
  chrome.tabs.goBack();
});

// Go home button
document.getElementById("goHome").addEventListener("click", () => {
  window.location.href = "https://www.google.com";
});

// Proceed anyway button (with confirmation)
document.getElementById("proceedAnyway").addEventListener("click", () => {
  const confirmed = confirm(
    "WARNING: You are about to visit a site that has been identified as a phishing attempt.\n\n" +
      "This site may:\n" +
      "• Steal your passwords and personal information\n" +
      "• Install malware on your device\n" +
      "• Compromise your financial accounts\n\n" +
      "Are you absolutely sure you want to proceed?"
  );

  if (confirmed) {
    const doubleConfirm = confirm(
      "FINAL WARNING: This action is NOT recommended.\n\n" +
        "Do you REALLY want to proceed to this dangerous site?"
    );

    if (doubleConfirm) {
      window.location.href = decodeURIComponent(blockedUrl);
    }
  }
});

const playSiren = () => {
  const sirenSound = new Audio("./lib/sirens.mp3");
  sirenSound.loop = true;
  sirenSound.volume = 0.7;
  sirenSound.play();
};

document.addEventListener("DOMContentLoaded", () => {
  playSiren();
});
