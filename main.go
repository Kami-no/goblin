package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"strconv"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const port int = 8080

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world! Welcome %s!\n", r.URL.Path[1:])
}

func info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Path:  %v\n", r.RequestURI)
	fmt.Fprintf(w, "Proto: %v\n", r.Proto)
}

func slow(w http.ResponseWriter, r *http.Request) {
	delaySlice, present := r.URL.Query()["delay"]
	if !present || len(delaySlice) == 0 {
		fmt.Fprintf(w, "Delay is not set\n")
		return
	}
	delay, err := strconv.Atoi(delaySlice[0])
	if err != nil {
		fmt.Fprintf(w, "Failed to define delay: %v!\n", err)
		return
	}
	fmt.Fprintf(w, "Delay (seconds): %v\n", delay)
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Fprintf(w, "Success\n")
}

func main() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatal("!ok")
	}
	fmt.Println(bi.Main.Version)
	// Router
	h2s := &http2.Server{}
	handler := http.NewServeMux()

	// Handlers
	handler.HandleFunc("/slow", slow)
	handler.HandleFunc("/info", info)
	handler.HandleFunc("/", hello)

	// m := &autocert.Manager{
	// 	Cache:      autocert.DirCache("golang-autocert"),
	// 	Prompt:     autocert.AcceptTOS,
	// 	HostPolicy: autocert.HostWhitelist("example.org", "www.example.org"),
	// }

	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", port),
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 180 * time.Second,
		IdleTimeout:  240 * time.Second,
		Handler:      h2c.NewHandler(handler, h2s),
		// TLSConfig:    m.TLSConfig(),
	}

	go func() {
		fmt.Printf("Running server at port %v...\n", port)
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
