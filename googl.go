package googl

import (
	"github.com/olebedev/go-googl/client"
	"github.com/olebedev/go-googl/client/url"
	"github.com/olebedev/go-googl/models"
)

// ShortenWithKey is a shortcut to operate from a server easier.
func ShortenWithKey(key, longURL string) (string, error) {
	sh := client.NewHTTPClient(nil)
	r, err := sh.URL.UrlshortenerURLInsert(
		url.NewUrlshortenerURLInsertParams().
			WithKey(&key).
			WithBody(
				&models.URL{
					LongURL: longURL,
				},
			), nil)
	if err != nil {
		return "", err
	}
	return r.Payload.ID, nil
}

// ShortenWithKeyFunc returnsis a shortcut to operate from a server easier.
func ShortenWithKeyFunc(key string) func(string) (string, error) {
	sh := client.NewHTTPClient(nil)
	return func(longURL string) (string, error) {
		r, err := sh.URL.UrlshortenerURLInsert(
			url.NewUrlshortenerURLInsertParams().
				WithKey(&key).
				WithBody(
					&models.URL{
						LongURL: longURL,
					},
				), nil)
		if err != nil {
			return "", err
		}
		return r.Payload.ID, nil
	}
}
