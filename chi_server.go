package main

import (
	"fmt"
	"io"
	"net/http"
	"github.com/go-chi/chi/v5"
)

// Chi framework server
func setupChiServer() {
	r := chi.NewRouter()

	// Standard GET
	r.Get("/api/v5/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"users": [{"id": 1, "name": "Eve"}]}`))
	})

	// Standard POST
	r.Post("/api/v5/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "User created"}`))
	})

	// Multiple methods on same endpoint using Route
	r.Route("/api/v5/products", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"products": [{"id": 1, "name": "Item"}]}`))
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"message": "Product created"}`))
		})
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "Product updated"}`))
		})
		r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "Product deleted"}`))
		})
	})

	// Non-standard method using MethodFunc
	r.MethodFunc("FUNKYTOWN", "/api/v5/funkytown", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Chi Funkytown!"}`))
	})

	// Another non-standard method
	r.MethodFunc("DANCE", "/api/v5/dance", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Chi Dance!"}`))
	})

	// Multiple non-standard methods
	r.MethodFunc("FUNKYTOWN", "/api/v5/custom", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"method": "funkytown", "framework": "chi"}`))
	})
	r.MethodFunc("PARTY", "/api/v5/custom", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"method": "party", "framework": "chi"}`))
	})

	// Discouraged: Missing method check - using HandleFunc
	r.HandleFunc("/api/v5/bad/no-method-check", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Accepts any method - bad practice!"}`))
	})

	// Discouraged: Missing error handling
	r.Post("/api/v5/bad/no-error-handling", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// No error handling for body read
		body, _ := io.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})

	// Discouraged: Missing input validation
	r.Post("/api/v5/bad/no-validation", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// No validation
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "No validation performed"}`))
	})

	// Discouraged: Missing Content-Type in some responses
	r.Get("/api/v5/bad/no-content-type", func(w http.ResponseWriter, r *http.Request) {
		// Missing Content-Type header
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data": "missing content-type"}`))
	})

	fmt.Println("Chi server running on :8084")
	http.ListenAndServe(":8084", r)
}

