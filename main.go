package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	ut "github.com/OluOwolabi/go-start/src/utils"
)

func main() {
	slog.Info("Starting the server...", "version", "1.0.0", "ulid", ut.GenerateUlid())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.Error("Server failed to start", "error", err)
		log.Fatal(err)
	}

}
