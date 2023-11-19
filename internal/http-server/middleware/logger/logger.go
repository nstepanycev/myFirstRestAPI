package logger

import (
	"net/http"
	"time"
	"golang.org/x/exp/slog"
	"github.com/gin-gonic/gin"
)


func ErrorMiddleWare(log *slog.Logger) func(next *gin.Engine) http.Handler{
	return func(next *gin.Engine) http.Handler{
		log = log.With(
            slog.String("component", "middleware/logger"),
        )
        log.Info("logger middleware enabled")

		fn := func(w http.ResponseWriter, r *http.Request) {
            entry := log.With(
                slog.String("method", r.Method),
                slog.String("path", r.URL.Path),
                slog.String("remote_addr", r.RemoteAddr),
                slog.String("user_agent", r.UserAgent()),
            )
			t1 := time.Now()
            
            defer func() {
                entry.Info("request completed",
                    slog.String("duration", time.Since(t1).String()),
                )
            }()
            next.ServeHTTP(w, r)
        }
		return http.HandlerFunc(fn)
	}
}