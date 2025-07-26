let session = pl.create(1000);

document.getElementById('fileInput').addEventListener('change', function(e) {
    const file = e.target.files[0];
    if (!file) return;

    const reader = new FileReader();
    reader.onload = function(e) {
        const program = e.target.result;
        console.log(program);
        session.consult(program, {
            success: () => { document.getElementById('output').innerText = 'Programa cargado correctamente\n'; },
            error: (err) => { document.getElementById('output').innerText = 'Error al cargar el programa\n' + err; }
        });
    };
    reader.readAsText(file);
});

function runQuery() {
    const query = document.getElementById('queryInput').value;
    document.getElementById('output').innerText = '';
    console.log(query);
    session.query(query, {
        success: () => { 
            session.answers(x => {
                if (x === false) {
                  document.getElementById("output").value += "No hay más soluciones.\n";
                  return;
                }
                document.getElementById("output").value += pl.format_answer(x) + "\n";
              });
         },
        fail: () => { document.getElementById('output').innerText += 'No hay respuesta\n'; },
        error: () => { document.getElementById('output').innerText += 'Error al realizar la consulta\n'; }
    });
}