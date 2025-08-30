package main

import (
	"os/exec"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	// 1. Abrir Microsoft Word
	cmd := exec.Command("open", "/Applications/Microsoft Word.app")
	err := cmd.Start()
	if err != nil {
		panic("No se pudo abrir Microsoft Word: " + err.Error())
	}

	// 2. Esperar más tiempo para que se abra Microsoft Word
	time.Sleep(8 * time.Second)

	// 3. Abrir un documento nuevo (Cmd+N en macOS)
	robotgo.KeyTap("n", "cmd")
	time.Sleep(5 * time.Second)

	// 4. Escribir datos en el documento
	time.Sleep(2 * time.Second)
	robotgo.TypeStr("Mynor Joel Lombardo Molina Guevara")
	time.Sleep(1 * time.Second)
	robotgo.KeyTap("enter")
	time.Sleep(1 * time.Second)
	robotgo.TypeStr("201503392")

	// 5. Tomar captura de pantalla
	time.Sleep(3 * time.Second)
	screenshotPath := "/Users/mac/Developer/university/IA1_2025/2S/Clase5/captura_word.png"
	err = robotgo.SaveCapture(screenshotPath)
	if err != nil {
		panic("No se pudo tomar la captura de pantalla: " + err.Error())
	}

	// 6. Confirmar captura
	robotgo.KeyTap("enter")
	robotgo.TypeStr("Captura guardada en: " + screenshotPath)
}
