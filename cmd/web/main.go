package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	port := flag.Int("port", 8080, "specify server port")
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Shrtnr!")
	})

	slog.Info("server starting", "port", *port)

	http.ListenAndServe(fmt.Sprintf(":%d", *port), mux)
}
