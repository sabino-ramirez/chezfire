package server

import (
	// "context"
	"log"
	"net/http"
)

func (s *server) handleSearch() http.HandlerFunc {
	log.Println("handleSearch initiated")

	// ctx := context.Background()

	return func(w http.ResponseWriter, r *http.Request) {

	}
}
