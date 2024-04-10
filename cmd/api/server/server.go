package server

import (
	"fmt"
	"log"
	"net/http"
	"structure/internal/http/route"
)

func StartServer() {
	router := route.NewRouter("8080")

	log.Printf("Server started on port %s...\n", router.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", router.Port), router))
}
