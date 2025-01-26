package main

import (
	"fmt"
	"net/http"

	"github.com/alexandrevicenzi/go-sse"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const PORT = 11814
const STATIC_PATH = "./static"

func main() {
	// Set up router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// Set up SSE agent
	fmt.Println("[setup] Creating SSE agent...")
	sse_agent := sse.NewServer(nil)
	defer sse_agent.Shutdown()
	// Set up file server
	fmt.Printf("[setup] Creating file server at \"%s\"...\n", STATIC_PATH)
	fs := http.FileServer(http.Dir(STATIC_PATH))
	// Set up routes
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	// Listen & serve
	fmt.Printf("Listening on port %d...\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), r)
}
