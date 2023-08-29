package main

import (
    "fmt"
	"io"
    "os"
    "os/exec"
    "path/filepath"
)

func main() {
    // Leer argumentos de línea de comandos
    args := os.Args[1:]

    // Verificar que se haya proporcionado un argumento
    if len(args) == 0 {
        fmt.Println("Debe proporcionar el nombre del script a ejecutar")
        os.Exit(1)
    }

    // Obtener el nombre del script a ejecutar
    scriptName := args[0]
    scriptPath := filepath.Join("./pkg/", scriptName, "script.go")

    // Verificar si el script existe
    if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
        fmt.Println("Script not found:", scriptPath)
        return
    }

	fmt.Println("Executing script:", scriptPath)
    // Ejecutar el script correspondiente
    cmd := exec.Command("go", "run", scriptPath)
    stdout, err := cmd.StdoutPipe()

    if err != nil {
        fmt.Println("Error al obtener la salida estándar del proceso:", err)
        os.Exit(1)
    }

    if err := cmd.Start(); err != nil {
        fmt.Println("Error al ejecutar el script:", err)
        os.Exit(1)
    }

    // Copiar la salida estándar del proceso secundario a la salida estándar del proceso principal
    if _, err := io.Copy(os.Stdout, stdout); err != nil {
        fmt.Println("Error al copiar la salida estándar del proceso:", err)
        os.Exit(1)
    }

    if err := cmd.Wait(); err != nil {
        fmt.Println("Error al esperar a que el proceso termine:", err)
        os.Exit(1)
    }
}