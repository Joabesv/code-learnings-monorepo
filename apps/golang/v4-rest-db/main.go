package main

import (
	"net/http"
	"v4-rest-db/pkg/logging"
	"v4-rest-db/routers"

	// load .env file
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func init() {
	logging.Setup()
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	logging.Info("Starting v4-rest-db server",
		zap.String("port", "5003"),
	)

	routes := routers.InitRouter()

	server := &http.Server{
		Addr:    ":5003",
		Handler: routes,
	}

	logging.Info("Server started successfully",
		zap.String("address", server.Addr),
	)

	if err := server.ListenAndServe(); err != nil {
		logging.Fatal("Failed to start server",
			zap.Error(err),
		)
	}
}
