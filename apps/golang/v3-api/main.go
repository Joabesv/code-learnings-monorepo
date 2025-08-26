package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func main() {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	apiRouter := router.Group("/api/v1")

	apiRouter.GET("/items", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pang",
			"data": []Item{
				{ID: 1, Description: "Item 1"},
				{ID: 2, Description: "Item 2"},
				{ID: 3, Description: "Item 3"},
			},
		})
	})

	router.Run()
}
