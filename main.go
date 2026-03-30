package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	port := "9000"

	// Prépare le serveur pour lire les fichiers du dossier "static"
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", http.StripPrefix("/", fs))

	// Détermine l'adresse IP locale de la machine
	ip := "127.0.0.1"
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		// Filtre pour trouver une adresse IPv4 valide et non locale (LAN)
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ip = ipnet.IP.String()
			break
		}
	}

	// Affiche l'URL d'accès et lance le serveur
	fmt.Printf("Serveur fonctionnant sur http://%s:%s/\n", ip, port)
	http.ListenAndServe(":"+port, nil)
}
