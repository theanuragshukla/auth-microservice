package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ContextKey string

const ContextKeyTraceID ContextKey = "traceID"

func AssignTraceID(ctx context.Context) context.Context {
	reqID := uuid.New()
	return context.WithValue(ctx, ContextKeyTraceID, reqID.String())
}

func GetTraceID(r *http.Request) string {
	ctx := r.Context()
	reqID := ctx.Value(ContextKeyTraceID)
	if ret, ok := reqID.(string); ok {
		return ret
	}
	return ""
}

func ReqIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		r = r.WithContext(AssignTraceID(ctx))
		next.ServeHTTP(w, r)
	})
}
