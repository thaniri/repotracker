package main

import (
	"github.com/thaniri/repotracker-app/rpController"
	"github.com/thaniri/repotracker-app/rpLogger"
	"net/http"
	"time"
)

func main() {
	defer rpLogger.Logger.Flush()

	handler := rpController.New()

	webApp := &http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	webApp.ListenAndServe()
}
