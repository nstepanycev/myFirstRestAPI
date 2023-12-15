package logger

import (
	// "net/http"
	// "time"
	"fmt"
    "context"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

type errorResponse struct {
	Message string `json:"message"`
}


func NewMiddleware() gin.HandlerFunc{
    return func(c *gin.Context){
        ctx := context.WithValue(c.Request.Context(),"ContextKey",c)
        c.Request = c.Request.WithContext(ctx)
        
        log := slog.With(
            slog.String("component", "middleware/logger"),
        )
        log.Info("logger midleware anable")
        c.Next()
        fmt.Println("Hello Dimmy!!")
    }
}

func middleWareReportResp(c *gin.Context, code int, message string){
    slog.Info(message)
    c.AbortWithStatusJSON(code, errorResponse{message})
}