package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	// 1. Abrir el diálogo "Ejecutar" con Win+R
	robotgo.KeyTap("r", "cmd")

	time.Sleep(500 * time.Millisecond)

	// 2. Escribir "notepad" y presionar Enter
	robotgo.TypeStr("notepad")
	robotgo.KeyTap("enter")

	// 3. Esperar a que se abra el Bloc de notas
	time.Sleep(1 * time.Second)

	// 4. Escribir texto
	robotgo.TypeStr("Hola! Esto fue escrito solo con robotgo")
	robotgo.KeyTap("enter")
	robotgo.TypeStr("Todo controlado por Go")
}
