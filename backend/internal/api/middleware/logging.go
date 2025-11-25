package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

// NewLogging returns a new Zap-based logging middleware.
func NewLogging(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			next.ServeHTTP(ww, r)

			logger.Info("request completed",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.String("remote_addr", r.RemoteAddr),
				zap.String("user_agent", r.UserAgent()),
				zap.Int("status", ww.Status()),
				zap.Int("bytes_written", ww.BytesWritten()),
				zap.String("duration", time.Since(t1).String()),
				zap.String("request_id", middleware.GetReqID(r.Context())),
			)
		}
		return http.HandlerFunc(fn)
	}
}
