package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Crear servidor

	fmt.Println("El servidor esa coorriendo en el puerto 3000")
	fmt.Printf("Run Server: http://localhost:3000/")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
