package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

// ContextID is our type to retrieve our context
type ContextID int

// ID is the only ID we've defined
const ID ContextID = 0

// SetID ...
func SetID(start int64) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), ID, strconv.FormatInt(start, 10))
			fmt.Println(r.Context())
			start++
			r = r.WithContext(ctx)
			next(w, r)
		}
	}
}

// GetID ...
func GetID(ctx context.Context) string {
	if val, ok := ctx.Value(ID).(string); ok {
		return val
	}
	return ""
}
