package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func Logging(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(lrw, r)

		logger.Info("request", "method", r.Method, "path", r.URL.Path, "status", lrw.statusCode, "remote_ip", r.RemoteAddr, "duration_ms", time.Since(start).Milliseconds())
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
