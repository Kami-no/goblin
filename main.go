package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
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
	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	// Wait for an interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt a graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = server.Shutdown(ctx)
}
