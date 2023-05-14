package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handlers
func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("El metodo es + " + r.Method)
	fmt.Fprintln(rw, "Hola fifi")
}

func PaginaNF(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}

func Error(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Este es un error", http.StatusConflict)
}

func main() {

	//Router
	http.HandleFunc("/", Hola)
	http.HandleFunc("/page", PaginaNF)
	http.HandleFunc("/error", Error)

	//Crear servidor

	fmt.Println("El servidor esta coorriendo en el puerto 3000")
	fmt.Printf("Run Server: http://localhost:3000/")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
