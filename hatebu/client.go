package hatebu

import (
	"net/http"
	"os"

	"github.com/garyburd/go-oauth/oauth"
)

type client struct {
}

func (c *client) post(url string, form map[string][]string) (*http.Response, error) {
	consumerKey := os.Getenv("HATENA_CONSUMER_KEY")
	consumerSecret := os.Getenv("HATENA_CONSUMER_SECRET")
	token := os.Getenv("HATENA_ACCESS_TOKEN")
	secret := os.Getenv("HATENA_ACCESS_SECRET")

	oauthClient := &oauth.Client{
		Credentials: oauth.Credentials{
			Token:  consumerKey,
			Secret: consumerSecret,
		},
		TemporaryCredentialRequestURI: "https://www.hatena.com/oauth/initiate",
		ResourceOwnerAuthorizationURI: "https://www.hatena.com/oauth/authorize",
		TokenRequestURI:               "https://www.hatena.com/oauth/token",
	}

	accessToken := oauth.Credentials{
		Token:  token,
		Secret: secret,
	}

	return oauthClient.Post(nil, &accessToken, url, form)
}
