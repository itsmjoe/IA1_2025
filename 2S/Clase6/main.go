package main

import (
	"fmt"
	"math"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	// Abrir Paint
	robotgo.KeyTap("r", "cmd")
	time.Sleep(500 * time.Millisecond)
	robotgo.TypeStr("mspaint")
	robotgo.KeyTap("enter")
	time.Sleep(3 * time.Second)

	// Buscar PID de Paint
	fpid, err := robotgo.FindIds("Paint")
	if err == nil && len(fpid) > 0 {
		fmt.Println("PID de Paint:", fpid[0])

		// Activar y maximizar ventana
		robotgo.ActivePid(fpid[0])
		time.Sleep(500 * time.Millisecond)
		robotgo.KeyTap("up", "cmd") // maximizar
		time.Sleep(500 * time.Millisecond)
	}

	// Obtener tamaño de pantalla
	w, h := robotgo.GetScreenSize()
	print(w, h)

	// Posición inicial centrada en el lienzo
	startX := 900
	startY := 600
	robotgo.Move(startX, startY)
	time.Sleep(500 * time.Millisecond)

	robotgo.KeyDown("ctrl")
	robotgo.ScrollDir(10, "up")
	robotgo.KeyUp("ctrl")

	// Dibujar espiral
	for i := 0; i < 500; i++ {
		angle := float64(i) * 0.1
		radius := float64(i) * 0.5
		px := int(float64(startX) + radius*math.Cos(angle))
		py := int(float64(startY) + radius*math.Sin(angle))

		robotgo.Move(px, py)
		robotgo.Click("left")
		time.Sleep(5 * time.Millisecond)
	}
}
