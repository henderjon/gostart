package ctx

import (
	"net/http"
	"time"
)

type CookieParams struct {
	Name     string
	Value    string
	TTL      time.Duration
	Path     string
	Domain   string
	IsSecure bool
}

// NewCookie creates a new Cookie
func NewCookie(opts CookieParams) *http.Cookie {
	return &http.Cookie{
		Name:    opts.Name,
		Value:   opts.Value,
		Expires: time.Now().UTC().Add(opts.TTL),
		Path:    opts.Path,
		Domain:  opts.Domain,
		// @NOTE using a non-secure cookie will not overwrite a secure cookie
		Secure:   opts.IsSecure,
		HttpOnly: true,
	}
}
