const output = document.getElementById("output");

document.getElementById("fileInput").addEventListener("change", async function(event) {
  const file = event.target.files[0];
  if (!file) return;
  
  const text = await file.text();
  const res = await fetch("/load", {
    method: "POST",
    headers: { "Content-Type": "text/plain" },
    body: text
  });

  output.textContent = await res.text();
});

async function runQuery() {
  const query = document.getElementById("queryInput").value;
  const res = await fetch("/query", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ query })
  });
  const data = await res.json();
  output.textContent = data.result.join("\n");
}

async function addFact() {
  const fact = document.getElementById("factInput").value;
  const res = await fetch("/add", {
    method: "POST",
    headers: { "Content-Type": "text/plain" },
    body: fact
  });
  output.textContent = await res.text();
}

function downloadCode() {
  window.open("/download", "_blank");
}