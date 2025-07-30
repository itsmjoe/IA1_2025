package main

import (
	"fmt"
	"os"

	"github.com/ichiban/prolog"
)

func main() {
	// Crear nueva máquina Prolog
	p := prolog.New(os.Stdin, os.Stdout)

	// Carga el código de Prolog
	if err := p.Exec(`
        padre(juan, maria).
        padre(juan, jose).
        madre(ana, maria).
        madre(ana, jose).
        hermano(X, Y) :- padre(P, X), padre(P, Y), X \= Y.
    `); err != nil {
		fmt.Printf("Error al cargar el código Prolog: %v\n", err)
		return
	}

	// Consulta: ¿maria y jose son hermanos?
	q := `hermano(maria, jose).`
	solutions, err := p.Query(q)
	if err != nil {
		fmt.Printf("Error al consultar: %v\n", err)
		return
	}
	defer solutions.Close()

	// Imprime las soluciones encontradas
	if solutions.Next() {
		fmt.Println("Sí, son hermanos.")
	} else {
		fmt.Println("No, no son hermanos.")
	}

	// Verifica si ocurre un error al consultar
	if err := solutions.Err(); err != nil {
		fmt.Printf("Error al consultar: %v\n", err)
	}
}
