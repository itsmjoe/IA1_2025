package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	// 1. Abrir el diálogo "Ejecutar" con Win+R
	robotgo.KeyTap("r", "cmd")
	robotgo.Sleep(1) // 1 segundo

	// 2. Abrir Notepad
	robotgo.TypeStr("notepad")
	robotgo.KeyTap("enter")
	robotgo.Sleep(1)

	// 3. Escribir en Notepad
	robotgo.WriteAll("Hola! Esto fue escrito solo con robotgo\nTodo controlado por Go")
	robotgo.Sleep(1)
	robotgo.KeyTap("v", "ctrl")
	time.Sleep(1 * time.Second)
}
