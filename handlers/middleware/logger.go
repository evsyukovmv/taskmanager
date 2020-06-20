package middleware

import (
	"github.com/evsyukovmv/taskmanager/logger"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		ctxRqId, ok := r.Context().Value(middleware.RequestIDKey).(string)
		if !ok {
			ctxRqId = "undefined"
		}

		defer func(start time.Time) {
			logger.Info("Request",
				zap.String("requestId", ctxRqId),
				zap.String("path", r.URL.Path),
				zap.String("method", r.Method),
				zap.String("remote", r.RemoteAddr),
				zap.Duration("latency", time.Since(start)),
				zap.Int("status", ww.Status()))
		}(time.Now())


		next.ServeHTTP(ww, r)
	}
	return http.HandlerFunc(fn)
}