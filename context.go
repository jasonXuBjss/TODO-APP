package main

import (
	"context"
	"net/http"
	"github.com/google/uuid"
)

type uuidKey int

const key uuidKey = 1

func contextWithUUID(ctx context.Context,id string) context.Context {
    return context.WithValue(ctx, key, id)
}

func uuidFromContext(ctx context.Context) (string, bool) {
    u, ok := ctx.Value(key).(string)
    return u, ok
}

func Middleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()

		if uuidHeader := r.Header.Get("X-UUID"); uuidHeader != "" {
            ctx = contextWithUUID(ctx, uuidHeader)
        } else {
            ctx = contextWithUUID(ctx, uuid.New().String())
        }

        r = r.WithContext(ctx)
        h.ServeHTTP(w, r)
    })
}
