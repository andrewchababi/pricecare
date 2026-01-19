const selectedAnalyses = new Map(); // id -> name

function addToSelectedTagList(cardEl) {
  const id = cardEl.dataset.analysisId;
  const name = cardEl.dataset.analysisName || id;

  if (selectedAnalyses.has(id)) {
    console.log("analysis has already been added");
    return;
  }
  selectedAnalyses.set(id, name);

  const tagList = document.querySelector(".tag-list");

  const tagEl = document.createElement("li");
  tagEl.className = "analysis-tag";
  tagEl.dataset.analysisId = id;

  tagEl.innerHTML = `
    <span class="analysis-name">${escapeHtml(id)}</span>
    <button class="analysis-remove" type="button" aria-label="Remove ${escapeHtml(id)}">
      <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 20 20" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
        <line x1="5" y1="5" x2="15" y2="15"></line>
        <line x1="15" y1="5" x2="5" y2="15"></line>
      </svg>
    </button>
  `;
  const removeBtn = tagEl.querySelector(".analysis-remove");
  removeBtn.addEventListener("click", () => removeAnalysis(removeBtn));

  tagList.appendChild(tagEl);
  console.log(selectedAnalyses)

  updateSelectedCount();
}

function updateSelectedCount() {
  const countEl = document.getElementById("panelTestCount");

  const n = selectedAnalyses.size;
  countEl.textContent = `${n} analyses selected`;
}

function escapeHtml(str) {
  return String(str)
    .replaceAll("&", "&amp;")
    .replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;")
    .replaceAll('"', "&quot;")
    .replaceAll("'", "&#039;");
}

async function calculatePanel() {
  const testIds = Array.from(selectedAnalyses.keys());

  if (testIds.length === 0) {
    console.warn("No analyses selected");
    return;
  }

  try {
    const res = await fetch("/api/calculate", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        testIds: testIds
      })
    });

    if (!res.ok) {
      const err = await res.json();
      console.error("Calculation error:", err);
      return;
    }

    const data = await res.json();

    const priceEl = document.getElementById("panelPrice");
    if (priceEl && typeof data.total_price === "number") {
      priceEl.textContent = `$${data.total_price.toFixed(2)}`;
    }

  } catch (err) {
    console.error("Network error calculating panel:", err);
  }
}

function removeAnalysis(btnEl) {
  const tagEl = btnEl.parentElement;

  const id = tagEl.dataset.analysisId;

  selectedAnalyses.delete(id);
  tagEl.remove();
  updateSelectedCount();
}

document.addEventListener("DOMContentLoaded", () => {
  window.addToSelectedTagList = addToSelectedTagList;
  window.calculatePanel = calculatePanel;
  window.removeAnalysis = removeAnalysis;
});
