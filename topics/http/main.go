package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"time"
)

// A compact, runnable tutorial demonstrating common usages of the
// Go net/http package. Run with `go run main.go`.

func main() {
	fmt.Println("Go net/http package tutorial — quick examples")
	fmt.Println("----------------------------------------------")

	basicServerClient()
	fmt.Println()

	routingAndHandlers()
	fmt.Println()

	middleware()
	fmt.Println()

	clientExamples()
}

func basicServerClient() {
	fmt.Println("1) Basic server and client")
	// Create a simple handler
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, %s!", r.URL.Query().Get("name"))
	})

	// Use httptest.Server so the example is runnable without reserving a port
	srv := httptest.NewServer(h)
	defer srv.Close()

	// Make a client request to the test server
	resp, err := http.Get(srv.URL + "/?name=Gopher")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Status: %s\nResponse: %s\n", resp.Status, string(body))
}

func routingAndHandlers() {
	fmt.Println("2) Routing and handlers")
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from /hello")
	})

	mux.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Current time: %s", time.Now().Format(time.RFC3339))
	})

	// Test with httptest
	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	fmt.Println("Response:", w.Body.String())
}

func middleware() {
	fmt.Println("3) Middleware example")
	// Simple logging middleware
	loggingMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handled with middleware")
	})

	wrapped := loggingMiddleware(handler)

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	wrapped.ServeHTTP(w, req)
	fmt.Println("Response:", w.Body.String())
}

func clientExamples() {
	fmt.Println("4) HTTP client examples")
	client := &http.Client{Timeout: 10 * time.Second}

	// Simple GET request (using test server)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Client test response")
	}))
	defer srv.Close()

	resp, err := client.Get(srv.URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Client response:", string(body))
}

