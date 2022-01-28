package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world! Welcome %s!\n", r.URL.Path[1:])
}

func info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Path:  %v\n", r.RequestURI)
	fmt.Fprintf(w, "Proto: %v\n", r.Proto)
}

func main() {
	// Router
	h2s := &http2.Server{}
	handler := http.NewServeMux()

	// Handlers
	handler.HandleFunc("/info", info)
	handler.HandleFunc("/", hello)

	// m := &autocert.Manager{
	// 	Cache:      autocert.DirCache("golang-autocert"),
	// 	Prompt:     autocert.AcceptTOS,
	// 	HostPolicy: autocert.HostWhitelist("example.org", "www.example.org"),
	// }

	server := &http.Server{
		Addr: ":8080",
		// TLSConfig: m.TLSConfig(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      h2c.NewHandler(handler, h2s),
	}

	go func() {
		fmt.Println("Running server...")
		log.Fatal(server.ListenAndServe())
		// log.Fatal(server.ListenAndServeTLS("", ""))
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
