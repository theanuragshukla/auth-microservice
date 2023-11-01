package middlewares

import (
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//reqID := GetTraceID(r)
		//logger.Info("Recieved request", zap.String("traceId", reqID), zap.String("ip", r.RemoteAddr))
		next.ServeHTTP(w, r)
	})
}
