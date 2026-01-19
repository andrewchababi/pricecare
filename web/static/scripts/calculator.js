const selectedAnalyses = new Map(); // id -> name

function addToSelectedTagList(cardEl) {
  if (!cardEl) return;

  const id = cardEl.dataset.analysisId;
  const name = cardEl.dataset.analysisName || id;

  if (!id) return;

  // Prevent duplicates
  if (selectedAnalyses.has(id)) return;

  selectedAnalyses.set(id, name);

  const tagList = document.querySelector(".tag-list");
  if (!tagList) return;

  // Create tag element
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

  tagList.appendChild(tagEl);

  // Update summary count
  updateSelectedCount();
}

function updateSelectedCount() {
  const countEl = document.getElementById("panelTestCount");
  if (!countEl) return;

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

// Event delegation: remove tags even if they were added later
document.addEventListener("click", (e) => {
  const btn = e.target.closest(".analysis-remove");
  if (!btn) return;

  const tag = btn.closest(".analysis-tag");
  if (!tag) return;

  const id = tag.dataset.analysisId;
  if (id) selectedAnalyses.delete(id);

  tag.remove();
  updateSelectedCount();
});

document.addEventListener("DOMContentLoaded", () => {
    window.addToSelectedTagList = addToSelectedTagList;
    window.escapeHtml = escapeHtml;
});

