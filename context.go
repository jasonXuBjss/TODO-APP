package main

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

//attach the uuid to context
//get the uuid from the context by using the key
//middleware ensure every req there is a uuid in that context

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

// type Logger struct {}

// func (Logger) Log(ctx context.Context, message string) {
//     if uuid, ok := uuidFromContext(ctx); ok {
//         message = fmt.Sprintf("UUID: %s - %s", uuid, message)
//     }
//     // do logging
//     fmt.Println(message)
// }

// func Request(req *http.Request) *http.Request {
//     ctx := req.Context()
//     if uuid, ok := uuidFromContext(ctx); ok {
//         req.Header.Add("X-UUID", uuid)
//     }
//     return req
// }