package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/ichiban/prolog"
)

var (
	prologVM *prolog.Interpreter
	code     bytes.Buffer
	mutex    sync.Mutex
)

func initProlog() {
	prologVM = prolog.New(os.Stdin, os.Stdout)
	code.Reset()
}

func main() {
	initProlog()

	http.HandleFunc("/load", handleLoad)
	http.HandleFunc("/add", handleAddFact)
	http.HandleFunc("/query", handleQuery)
	http.HandleFunc("/download", handleDownload)

	fmt.Println("Servidor iniciado en http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

// POST /load - carga el código prolog desde un archivo de texto
func handleLoad(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	initProlog()
	code.Write(body)

	err = prologVM.Exec(code.String())
	if err != nil {
		http.Error(w, "Error al cargar el código Prolog", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Código Prolog cargado exitosamente"))
}

// POST /add - agrega un hecho al código Prolog
func handleAddFact(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	fact := string(body)
	code.WriteString("\n" + fact)

	err = prologVM.Exec(fact)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al agregar el hecho: %v", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Hecho agregado correctamente"))
}

// POST /query - ejecuta una consulta Prolog
func handleQuery(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var query struct {
		Query string `json:"query"`
	}
	if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
		http.Error(w, "Consulta inválida", http.StatusBadRequest)
		return
	}

	solutions, err := prologVM.Query(query.Query)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error en la consulta %v", err), http.StatusInternalServerError)
		return
	}
	defer solutions.Close()

	var results []string
	for solutions.Next() {
		var m map[string]any
		if err := solutions.Scan(&m); err != nil {
			http.Error(w, fmt.Sprintf("Error al escanear resultados: %v", err), http.StatusInternalServerError)
			return
		}
		results = append(results, fmt.Sprintf("%v", m))
	}

	if len(results) == 0 {
		results = append(results, "No se encontraron resultados")
	}

	resp := map[string]interface{}{
		"results": results,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// GET /download - descarga el código Prolog actual
func handleDownload(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	w.Header().Set("Content-Disposition", "attachment; filename=code.pl")
	w.Header().Set("Content-Type", "text/plain")
	w.Write(code.Bytes())
}
