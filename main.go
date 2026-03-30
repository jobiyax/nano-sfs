package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := "9000" // Port du serveur

	// Sert les fichiers du dossier "static" à la racine
	fs := http.FileServer(http.Dir("static"))

	// Retire le préfixe "/" pour pointer sur "static"
	http.Handle("/", http.StripPrefix("/", fs))

	fmt.Println("Serveur fonctionnant sur http://localhost:" + port + "/")
	http.ListenAndServe(":"+port, nil)
}
