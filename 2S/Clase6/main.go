package main

import (
	"fmt"
	"math"
	"os/exec"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("Creando imagen grande para dibujar...")
	canvasPath := "/Users/mac/Downloads/lienzo_actividad6.png"

	cmd := exec.Command("sips", "-s", "format", "png", "-z", "800", "1200", "/System/Library/Desktop Pictures/Solid Colors/Silver.png", "--out", canvasPath)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error creando imagen grande: %v\n", err)
		// Método alternativo: crear imagen grande con sips
		cmd = exec.Command("sips", "-s", "format", "png", "-z", "800", "1200", "/Applications/Preview.app/Contents/Resources/Preview.icns", "--out", canvasPath)
		err = cmd.Run()
		if err != nil {
			panic("No se pudo crear la imagen: " + err.Error())
		}
	}

	// Verificar que el archivo se creó
	cmd = exec.Command("ls", "-la", canvasPath)
	output, err := cmd.Output()
	if err != nil {
		panic("Error verificando archivo: " + err.Error())
	}
	fmt.Printf("Archivo creado: %s\n", string(output))

	// Abrir Preview directamente con la imagen
	fmt.Println("Abriendo Preview...")
	cmd = exec.Command("open", canvasPath)
	err = cmd.Start()
	if err != nil {
		fmt.Printf("Error abriendo con 'open': %v\n", err)
		// Intentar con método alternativo
		cmd = exec.Command("open", "-a", "Preview", canvasPath)
		err = cmd.Start()
		if err != nil {
			panic("No se pudo abrir el archivo: " + err.Error())
		}
	}

	fmt.Println("Esperando a que Preview se abra...")
	time.Sleep(3 * time.Second)
	time.Sleep(5 * time.Second)

	// Intentar encontrar Preview varias veces
	var fpid []int
	for i := 0; i < 5; i++ {
		fpid, err = robotgo.FindIds("Preview")
		if err == nil && len(fpid) > 0 {
			fmt.Printf("¡Preview encontrado! PID: %d\n", fpid[0])
			break
		}
		fmt.Printf("Intento %d: Preview no encontrado, esperando...\n", i+1)
		time.Sleep(2 * time.Second)
	}

	if len(fpid) == 0 {
		fmt.Println("Preview no se abrió correctamente. Intentando abrir manualmente...")
		// Forzar apertura de Preview
		cmd = exec.Command("open", "-a", "Preview", canvasPath)
		cmd.Run()
		time.Sleep(3 * time.Second)

		fpid, err = robotgo.FindIds("Preview")
		if err != nil || len(fpid) == 0 {
			panic("No se pudo encontrar Preview después de múltiples intentos")
		}
	}

	// Activar ventana de Preview
	fmt.Println("Activando ventana de Preview...")
	robotgo.ActivePid(fpid[0])
	time.Sleep(2 * time.Second)

	robotgo.ActivePid(fpid[0])
	time.Sleep(1 * time.Second)

	// Activar herramientas de markup/anotación usando el menú
	fmt.Println("Activando herramientas de anotación mediante menú...")

	// Intentar acceso directo al menú Tools
	robotgo.KeyTap("t", "cmd", "alt") // Menú Tools
	time.Sleep(1 * time.Second)

	// También intentar con Show Markup Toolbar directamente
	fmt.Println("Intentando mostrar barra de herramientas de marcado...")
	robotgo.KeyTap("a", "cmd", "shift") // Show Markup Toolbar
	time.Sleep(2 * time.Second)

	// Método alternativo: usar el menú View
	robotgo.KeyTap("v", "cmd") // Menú View
	time.Sleep(1 * time.Second)
	robotgo.TypeStr("Show Markup Toolbar") // Escribir la opción
	time.Sleep(1 * time.Second)
	robotgo.KeyTap("enter") // Seleccionar
	time.Sleep(2 * time.Second)

	// Escape para cerrar cualquier menú abierto
	robotgo.KeyTap("escape")
	time.Sleep(1 * time.Second) // Obtener información de pantalla
	w, h := robotgo.GetScreenSize()
	fmt.Printf("Tamaño de pantalla: %dx%d\n", w, h)

	// Calcular posición central más conservadora
	startX := w / 2
	startY := h / 2
	fmt.Printf("Posición para dibujar: %d,%d\n", startX, startY)

	// Mover mouse al centro de la pantalla
	robotgo.Move(startX, startY)
	time.Sleep(1 * time.Second)

	// Hacer clic para asegurar que la ventana está activa
	robotgo.Click("left")
	time.Sleep(1 * time.Second)

	// Método más específico para seleccionar herramienta de dibujo
	fmt.Println("Buscando y seleccionando herramienta de sketch/pencil...")

	// Una vez que se abre la barra de herramientas, buscar el ícono de lápiz/sketch
	// Las herramientas suelen estar en la parte superior de la ventana
	toolbarY := 120 // Altura típica de la barra de herramientas en Preview

	// Buscar en diferentes posiciones típicas del ícono de sketch/pencil
	fmt.Println("Intentando seleccionar herramienta de sketch...")

	// Posición típica 1: lado izquierdo de la barra
	fmt.Println("Probando posición 1 (izquierda)...")
	robotgo.Move(startX-200, toolbarY)
	time.Sleep(500 * time.Millisecond)
	robotgo.Click("left")
	time.Sleep(1 * time.Second)

	// Posición típica 2: centro-izquierda
	fmt.Println("Probando posición 2 (centro-izquierda)...")
	robotgo.Move(startX-100, toolbarY)
	time.Sleep(500 * time.Millisecond)
	robotgo.Click("left")
	time.Sleep(1 * time.Second)

	// Posición típica 3: centro
	fmt.Println("Probando posición 3 (centro)...")
	robotgo.Move(startX, toolbarY)
	time.Sleep(500 * time.Millisecond)
	robotgo.Click("left")
	time.Sleep(1 * time.Second)

	// Intentar también con atajos de teclado específicos para sketch
	fmt.Println("Probando atajos de teclado para sketch...")
	robotgo.KeyTap("s", "cmd") // Sketch tool
	time.Sleep(1 * time.Second)

	// Atajo alternativo para pencil/draw
	robotgo.KeyTap("p", "cmd") // Pencil tool
	time.Sleep(1 * time.Second)

	// Verificar que una herramienta está seleccionada haciendo un pequeño movimiento de prueba
	fmt.Println("Verificando selección de herramienta...")
	currentX, currentY := robotgo.Location()
	robotgo.Move(currentX+5, currentY+5)
	robotgo.MouseDown("left")
	time.Sleep(100 * time.Millisecond)
	robotgo.MouseUp("left")
	time.Sleep(500 * time.Millisecond) // Mover de nuevo al centro para dibujar
	robotgo.Move(startX, startY)
	time.Sleep(1 * time.Second)

	// Pausa adicional para asegurar que las herramientas están activas
	fmt.Println("Esperando a que las herramientas se activen...")
	time.Sleep(3 * time.Second)

	// Dibujar círculo con mejor visibilidad
	fmt.Println("Iniciando dibujo del círculo...")
	fmt.Println("OBSERVA la pantalla - si ves el cursor moviéndose, el dibujo está funcionando...")

	// Hacer clic inicial para comenzar el trazo
	robotgo.MouseDown("left")

	// Dibujar círculo más grande para la imagen ampliada
	radius := 150.0 // Radio más grande para imagen de 1200x800
	steps := 72     // Menos pasos pero más grandes para mejor visibilidad

	for i := 0; i <= steps; i++ {
		angle := float64(i) * 2 * math.Pi / float64(steps) // Círculo completo
		px := int(float64(startX) + radius*math.Cos(angle))
		py := int(float64(startY) + radius*math.Sin(angle))

		robotgo.Move(px, py)
		time.Sleep(20 * time.Millisecond) // Más lento para asegurar el trazo

		// Mostrar progreso cada 10 pasos
		if i%10 == 0 {
			fmt.Printf("Progreso: %d%%\n", (i*100)/steps)
		}
	}

	robotgo.MouseUp("left")
	fmt.Println("¡Círculo completado!")

	// Guardar como nuevo archivo con nombre específico
	time.Sleep(2 * time.Second)
	fmt.Println("Guardando archivo como actividad6_circulo.png...")

	// Usar "Guardar como" para cambiar el nombre
	robotgo.KeyTap("s", "cmd", "shift") // Save As
	time.Sleep(2 * time.Second)

	// Escribir el nuevo nombre del archivo
	robotgo.TypeStr("actividad6_circulo.png")
	time.Sleep(1 * time.Second)

	// Presionar Enter para guardar
	robotgo.KeyTap("enter")
	time.Sleep(3 * time.Second)

	// Verificar que el nuevo archivo se guardó
	newCanvasPath := "/Users/mac/Downloads/actividad6_circulo.png"
	cmd = exec.Command("ls", "-la", newCanvasPath)
	output2, _ := cmd.Output()
	fmt.Printf("Archivo guardado como: %s\n", string(output2))

	fmt.Println("¡Proceso completado! Revisa el archivo:", newCanvasPath)
	fmt.Println("Si no se dibujó el círculo, puede que Preview no tenga herramientas de dibujo habilitadas.")
}
