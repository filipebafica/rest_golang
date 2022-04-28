package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/filipebafica/rest_golang/src/handlers"
)

func main () {
	l := log.New(os.Stdout, "products-api", log.LstdFlags)

	// creates an end point "object"
	productsHandler := handlers.NewProducts(l)

	// matches the request to the correspont handler function
	mux := http.NewServeMux()
	mux.Handle("/", productsHandler)

	// sets the server parameters
	server := &http.Server{
		Addr: ":9090",
		Handler: mux,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// light weighted thread (not OS thread) that will execute in concurrency mode
	go func (){
		err := server.ListenAndServe()
		if err != nil {
			l.Println(err)
		}
	}()

	// Signal handling
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <- sigChan
	l.Println("Recieved terminate, graceful shutdown\n", sig)

	// waits until all current requests finishes before a server shutdown happen
	ctx, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	server.Shutdown(ctx)
}
