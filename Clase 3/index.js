let session = pl.create(1000);
let prologCode = "";
    
document.getElementById('fileInput').addEventListener('change', function(event) {
  const file = event.target.files[0];
  if (!file) return;
  
  const reader = new FileReader();
  reader.onload = function(e) {
    prologCode = e.target.result;
    session = pl.create(1000);
    session.consult(prologCode, {
      success: function() { document.getElementById("output").value = "Código cargado correctamente.\n"; },
      error: function(err) { document.getElementById("output").value = "Error al cargar el código: " + err; }
    });
  };
  reader.readAsText(file);
});

function runQuery() {
  const query = document.getElementById("queryInput").value;
  document.getElementById("output").value = "Ejecutando consulta...\n";
  
  session.query(query, {
    success: function() {
      session.answers(x => {
        if (x === false) {
          document.getElementById("output").value += "No hay más soluciones.\n";
          return;
        }
        document.getElementById("output").value += pl.format_answer(x) + "\n";
      });
    },
    fail: function() {
      document.getElementById("output").value += "No hay resultados.\n";
    },
    error: function(err) {
      document.getElementById("output").value += "Error en la consulta: " + err + "\n";
    }
  });
}

function addFact() {
  const newFact = document.getElementById("factInput").value.trim();
  if (newFact === "") {
    alert("Por favor, ingrese un hecho.");
    return;
  }

  prologCode += newFact + "\n";

  session = pl.create(1000);
  session.consult(prologCode, {
    success: function() {
      document.getElementById("output").value = "Hecho agregado correctamente.\n";
    },
    error: function(err) {
      document.getElementById("output").value = "Error al agregar el hecho: " + err;
    }
  });

  document.getElementById("factInput").value = "";
}

function downloadPrologFile() {
  const element = document.createElement('a');
  element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(prologCode));
  element.setAttribute('download', 'code.pl');
  element.style.display = 'none';
  document.body.appendChild(element);
  element.click();
  document.body.removeChild(element);
}