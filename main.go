package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world! Welcome %s!\n", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)

	// m := &autocert.Manager{
	// 	Cache:      autocert.DirCache("golang-autocert"),
	// 	Prompt:     autocert.AcceptTOS,
	// 	HostPolicy: autocert.HostWhitelist("example.org", "www.example.org"),
	// }
	server := &http.Server{
		Addr: ":8080",
		// TLSConfig: m.TLSConfig(),
	}
	// log.Fatal(server.ListenAndServeTLS("", ""))

	log.Fatal(server.ListenAndServe())
}
