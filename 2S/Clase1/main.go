package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/ichiban/prolog"
)

func main() {
	// Crear nueva máquina Prolog
	p := prolog.New(nil, nil)

	// Código Prolog (como en Tau Prolog)
	program := `
        padre(juan, maria).
        padre(juan, jose).
        madre(ana, maria).
        madre(ana, jose).

        hermano(X, Y) :- padre(P, X), padre(P, Y), X \= Y.
    `

	// Ejecutar el código Prolog
	if err := p.Exec(program); err != nil {
		log.Fatalf("Error cargando el programa: %v", err)
	}

	// Consulta: ¿maria y jose son hermanos?
	q := `hermano(maria, jose).`
	solutions, err := p.Query(strings.NewReader(q))
	if err != nil {
		log.Fatalf("Error ejecutando la consulta: %v", err)
	}

	// Verificar si hay al menos una solución
	if solutions.Next() {
		fmt.Println("✅ Sí, son hermanos.")
		var s engine.Struct
		if err := solutions.Scan(&s); err == nil {
			fmt.Printf("Resultado: %+v\n", s)
		}
	} else {
		fmt.Println("❌ No son hermanos.")
	}

	// Verificar si hubo errores durante la iteración
	if err := solutions.Err(); err != nil {
		log.Fatalf("Error obteniendo resultados: %v", err)
	}
}
