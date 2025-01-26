package main

import (
	"fmt"
	"net/http"

	"github.com/alexandrevicenzi/go-sse"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const PORT = 11814

func main() {
	// Set up router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// Set up SSE agent
	sse_agent := sse.NewServer(nil)
	defer sse_agent.Shutdown()
	// Set up file server
	fs := http.FileServer(http.Dir("./static"))
	// Set up routes
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	// Listen & serve
	fmt.Printf("Listening on port %d...\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), r)
}
