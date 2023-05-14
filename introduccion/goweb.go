package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handlers
func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hola Mundo")
}

func main() {

	//Router
	http.HandleFunc("/", Hola)

	//Crear servidor

	fmt.Println("El servidor esta coorriendo en el puerto 3000")
	fmt.Printf("Run Server: http://localhost:3000/")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
