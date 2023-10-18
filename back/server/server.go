package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sabino-ramirez/chezfire/auth"
	"github.com/sabino-ramirez/chezfire/services/spotify"
)

// pass server dependencies here
type deps struct {
	authClient *spotify.Client
	authHelper *auth.AuthTokenHelper
}

type server struct {
	router chi.Router
	client *spotify.Client // standard clientCredentials

	// for auth token flow
	deps
}

// constructor function for server
func NewServer(httpClient *http.Client) *server {
	s := &server{
		router: chi.NewRouter(),
		client: spotify.New(httpClient),
	}

	s.routes()
	return s
}

// Server needs to implement this in order to be used as a handler
// in ListenAndServe in main.go
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// calling this function in the constructor initializes the routes
func (s *server) routes() {

	currTok, _ := s.client.Token()
	log.Println("start up token:", currTok)

	// initialize routes
	s.router.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte("Hello World"))
		if err != nil {
			log.Println(err)
		}
	})
	s.router.Get("/login", s.handleLogin())
	s.router.Get("/callback", s.handleCallback())
	s.router.Get("/featuredPlaylists", s.handleFeaturedPlaylists())
	s.router.Get("/search", s.handleSearch())
}
