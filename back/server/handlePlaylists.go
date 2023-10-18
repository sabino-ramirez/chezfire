package server

import (
	"context"
	"log"
	"net/http"
)

func (s *server) handleFeaturedPlaylists() http.HandlerFunc {
	log.Println("handleFeaturedPlaylists called")

	ctx := context.Background()

	return func(_ http.ResponseWriter, _ *http.Request) {
		msg, page, err := s.client.FeaturedPlaylists(ctx)

		if err != nil {
			log.Println("featured playlists err:", err)
		}

		log.Println(msg)

		for _, playlist := range page.Playlists {
			log.Println(" ", playlist.Name)
		}

		currTok2, _ := s.client.Token()
		log.Println("info was retrieved with token: \n", currTok2)
	}
}
