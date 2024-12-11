package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Servir archivos estáticos desde el directorio "static"
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Servidor iniciado en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
