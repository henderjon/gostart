package ctx

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func getMockLocalSessionPingRequest(params CookieParams) (*http.Request, error) {

	u, err := url.ParseRequestURI("http://localhost.com/")
	if err != nil {
		return nil, errors.New("unable to create mock request")
	}

	req := httptest.NewRequest("GET", u.String(), nil)

	params.Value = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsYWEiOjE1MjE1NTA5ODcsInRhbCI6MzAyLCJleHAiOjIwMjE2NTg5ODcsImp0aSI6IjJkYmY5YmNhLWYyZjItNGY4MS04YjU3LTA3MjM1YmFjMTg4YSIsImlhdCI6MTUyMTU0ODY4Mn0.Lvo4Umv0q97ZCIyuqPrrhOOH8FAglfsGn2C6Ce4exA0`
	req.AddCookie(NewCookie(params))

	return req, nil
}

func TestLocalSessionParser(t *testing.T) {

	salt := `seasalt`
	params := CookieParams{
		Name:     "local",
		TTL:      1800,
		Path:     "/",
		Domain:   ".localhost.com",
		IsSecure: false,
	}

	req, err := getMockLocalSessionPingRequest(params)
	if err != nil {
		t.Error(err)
	}

	parser := NewLocalSessionParser(params.Name, salt, params.TTL)

	// contexts are contextual ... #sigh
	localSession, err := parser(req)
	if err != nil {
		t.Error("error parsing local session")
	}

	if localSession.JWTID != "2dbf9bca-f2f2-4f81-8b57-07235bac188a" {
		t.Error("error unmarshaling JWT")
	}
}
