package main

import (
	"context"
	"net/http"
)

type ctxTmpl struct {
	Val int
}

// custom type to use as a key to avoid collisions within the context
type ctxKeyTmpl int

// initialize a value of the custom type to avoid collisions within the context
var ctxTmplKey ctxKeyTmpl

// function to set a *ctxTmpl at our custom typed key
func ctxSetTmpl(ctx context.Context, b *ctxTmpl) context.Context {
	return context.WithValue(ctx, ctxTmplKey, b)
}

// func type to use as psuedo-interface in external applications
type ctxTmplGetter func(ctx context.Context) (*ctxTmpl, bool)

// function to get a *ctxTmpl at our custom typed key
func ctxGetTmpl(ctx context.Context) (*ctxTmpl, bool) {
	b, ok := ctx.Value(ctxTmplKey).(*ctxTmpl)
	return b, ok
}

func ctxTmplHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = ctxSetTmpl(ctx, &ctxTmpl{})
		*r = *r.WithContext(ctx)
	})
}
