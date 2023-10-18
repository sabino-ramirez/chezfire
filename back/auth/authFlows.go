package auth

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type ClientCredHelper struct {
	// should be pointer because we will update the fields
	Config *clientcredentials.Config
}

type AuthTokenHelper struct {
	config *oauth2.Config
}

type AuthTokenOption func(cch *AuthTokenHelper)

// example option
func AddScopes(scopes ...string) AuthTokenOption {
	return func(ath *AuthTokenHelper) {
		ath.config.Scopes = scopes
	}
}

func AddRedirectURL(url string) AuthTokenOption {
	return func(ath *AuthTokenHelper) {
		ath.config.RedirectURL = url
	}
}

func NewClientCredHelper() *ClientCredHelper {
	clientCredsConfig := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     TokenURL,
		AuthStyle:    2, // AuthStyleHeader: id and secret sent in header
		EndpointParams: url.Values{
			"grant_type": {GrantType},
		},
	}

	// make clientCredHelper with configs
	cch := &ClientCredHelper{clientCredsConfig}

	return cch
}

func NewAuthTokenHelper(options ...AuthTokenOption) *AuthTokenHelper {
	authTokenConfig := &oauth2.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:   os.Getenv("AUTH_URL"),
			TokenURL:  os.Getenv("TOKEN_URL"),
			AuthStyle: 0,
		},
	}

	ath := &AuthTokenHelper{authTokenConfig}

	// if any options were passed, add them to config
	for _, option := range options {
		option(ath)
	}

	return ath
}

func (ath AuthTokenHelper) AuthURL(state string, opts ...oauth2.AuthCodeOption) string {
	return ath.config.AuthCodeURL(state, opts...)
}

func (ath AuthTokenHelper) Token(
	ctx context.Context,
	state string,
	r *http.Request,
	opts ...oauth2.AuthCodeOption,
) (*oauth2.Token, error) {
	values := r.URL.Query()

	if err := values.Get("error"); err != "" {
		return nil, errors.New("spotify: auth failed: " + err)
	}

	code := values.Get("code")
	if code == "" {
		return nil, errors.New("spotify: no access code received: ")
	}

	trueState := values.Get("state")
	if trueState != state {
		return nil, errors.New("spotify: state doesn't match")
	}

	return ath.config.Exchange(ctx, code, opts...)
}

func (ath AuthTokenHelper) Client(ctx context.Context, token *oauth2.Token) *http.Client {
	return ath.config.Client(ctx, token)
}

// type tokenSource struct {
// 	ctx  context.Context
// 	conf *clientcredentials.Config
// }
//
// // just call underlying func for now
// func (cch ClientCredHelper) TokenSource(ctx context.Context) oauth2.TokenSource {
// 	source := cch.config.TokenSource(ctx)
//
// 	return source
// }
//
// func (cch ClientCredHelper) Client(ctx context.Context) *http.Client {
// 	return cch.config.Client(ctx)
// }
