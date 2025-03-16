package middlewares

import (
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var logger, _ = zap.NewProduction()

func LoggingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        logger.Info("Incoming request",
            zap.String("method", c.Request.Method),
            zap.String("path", c.Request.URL.Path),
        )
        c.Next()
    }
}
