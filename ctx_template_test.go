package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCtxTmplHandler(t *testing.T) {

	u, err := url.ParseRequestURI("http://localhost.com/a/path")
	if err != nil {
		t.Error("unable to create mock request")
	}

	req := httptest.NewRequest("GET", u.String(), nil)

	w := httptest.NewRecorder()
	ctxTmplHandler().ServeHTTP(w, req)

	// contexts are contextual ... #sigh
	ctx, ok := ctxGetTmpl(req.Context())
	if !ok {
		t.Error("context not found")
	}

	expected := 0

	if diff := cmp.Diff(ctx.Val, expected); diff != "" {
		t.Errorf("context not found: (-got +want)\n%s", diff)
	}

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Error("failed to get status:", http.StatusOK)
	}

}
