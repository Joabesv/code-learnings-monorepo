package middleware

import (
	"context"
	"net/http"
	"time"

	"v4-rest-db/pkg/logging"

	"go.uber.org/zap"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int
}

func (lrw *loggingResponseWriter) WriteHeader(statusCode int) {
	lrw.statusCode = statusCode
	lrw.ResponseWriter.WriteHeader(statusCode)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := lrw.ResponseWriter.Write(b)
	lrw.size += size
	return size, err
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		correlationID := generateCorrelationID()

		logger := logging.L().With(
			zap.String("correlation_id", correlationID),
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
			zap.String("user_agent", r.UserAgent()),
		)

		w.Header().Add("X-Correlation-ID", correlationID)

		lrw := &loggingResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		ctx := context.WithValue(r.Context(), "logger", logger)
		r = r.WithContext(ctx)

		logger.Debug("Request started")

		next.ServeHTTP(lrw, r)

		duration := time.Since(start)

		logger.Info("Request completed",
			zap.Int("status_code", lrw.statusCode),
			zap.Int("response_size", lrw.size),
			zap.Duration("duration", duration),
		)
	})
}

func GetLoggerFromContext(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value("logger").(*zap.Logger); ok {
		return logger
	}
	return logging.L()
}

func generateCorrelationID() string {
	return time.Now().Format("20060102150405") + "-" + randomString(8)
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}
