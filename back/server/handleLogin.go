package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sabino-ramirez/chezfire/auth"
	"github.com/sabino-ramirez/chezfire/services/spotify"
)

var (
	state = "abc123"
	ch    = make(chan *spotify.Client)
)

func (s *server) handleLogin() http.HandlerFunc {
	log.Println("handleLogin invoked")

	return func(w http.ResponseWriter, _ *http.Request) {
		authHelper := auth.NewAuthTokenHelper(auth.AddRedirectURL(os.Getenv("REDIRECT_URL")))
		s.authHelper = authHelper

		url := authHelper.AuthURL(state)

		log.Println("log in at:", url)
		fmt.Fprintf(w, "log in at: %v", url)

		client := <-ch

		s.authClient = client

		tok, _ := client.Token()
		log.Println("token:", tok)
	}
}

func (s *server) handleCallback() http.HandlerFunc {
	log.Println("handleCallback invoked")

	return func(w http.ResponseWriter, r *http.Request) {
		token, err := s.authHelper.Token(r.Context(), state, r)

		if err != nil {
			http.Error(w, "Couldn't get token", http.StatusForbidden)
			log.Fatal(err)
		}

		if st := r.FormValue("state"); st != state {
			http.NotFound(w, r)
			log.Fatalf("State mismatch: %s != %s\n", st, state)
		}

		client := spotify.New(s.authHelper.Client(r.Context(), token))

		fmt.Fprint(w, "Login completed")

		ch <- client
	}
}
