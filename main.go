package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func main() {
	// Handler qui répond à toutes les requêtes avec "Hello World"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	// Démarrer le serveur sur le port 8080 et gérer les erreurs
	fmt.Printf("Compiled with Go version %s\n", runtime.Version())
	fmt.Println("Server started on the 8080 port...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}