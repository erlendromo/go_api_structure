package route

import (
	"net/http"
	"structure/internal/http/handlers"
)

type Router struct {
	Port string
}

func NewRouter(p string) *Router {
	return &Router{
		Port: p,
	}
}

func (route *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handlers.HandleRoot(w, r)
	default:
		http.Error(w, "endpoint not found", 404)
	}
}
