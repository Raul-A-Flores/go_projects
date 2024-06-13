package main

import (
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("userID")
		w.Write([]byte("User ID: " + userID))
	})

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	middlewareChain := MiddlewareChain(RequestLoggerMiddleware, RequireAuthMiddleWare)

	server := http.Server{
		Addr: s.addr,
		// Handler: RequireAuthMiddleWare(RequestLoggerMiddleware(router)),
		Handler: middlewareChain(router),
	}

	log.Printf("Sever has started on port %s", s.addr)

	return server.ListenAndServe()

}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method %s, path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

// Authentication

func RequireAuthMiddleWare(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// check if user is authenticatied

		token := r.Header.Get("Authorization")
		if token != "Bearer token" {

			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

type Middleware func(next http.Handler) http.HandlerFunc

func MiddlewareChain(middlewares ...Middleware) Middleware {

	return func(next http.Handler) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {

			next = middlewares[i](next)
		}

		return next.ServeHTTP

	}

}
