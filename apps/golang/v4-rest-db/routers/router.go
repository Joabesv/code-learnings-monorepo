package routers

import (
	"net/http"
	"os"
	"time"
	"v4-rest-db/pkg/logging"
	"v4-rest-db/pkg/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

func InitRouter() *gin.Engine {
	logging.Info("Initializing router")

	r := gin.New()

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		logging.Info("Route registered",
			zap.String("method", httpMethod),
			zap.String("path", absolutePath),
			zap.String("handler", handlerName),
			zap.Int("handlers_count", nuHandlers),
		)
	}

	r.Use(customGinLogger())
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		logger := middleware.GetLoggerFromContext(c.Request.Context())
		logger.Info("Root endpoint accessed")
		c.JSON(http.StatusOK, gin.H{"message": "ok swagger"})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/health", func(c *gin.Context) {
			start := time.Now()
			logger := middleware.GetLoggerFromContext(c.Request.Context())

			logger.Debug("Health check requested")

			response := gin.H{
				"message": "ok",
				"debug":   os.Getenv("DEBUG_ENV_VAR"),
				"time":    time.Now().Format(time.RFC3339),
			}

			c.JSON(http.StatusOK, response)

			logger.Info("Health check completed",
				zap.Duration("duration", time.Since(start)),
				zap.Int("status", 200),
			)
		})
	}

	logging.Info("Router initialized successfully")
	return r
}

func customGinLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		logger := logging.L().With(
			zap.String("method", param.Method),
			zap.String("path", param.Path),
			zap.Int("status", param.StatusCode),
			zap.String("ip", param.ClientIP),
			zap.String("user_agent", param.Request.UserAgent()),
			zap.Duration("latency", param.Latency),
			zap.Int("body_size", param.BodySize),
		)

		if param.StatusCode >= 400 {
			logger.Error("Request failed")
		} else {
			logger.Info("Request completed")
		}

		return ""
	})
}
