package routes

import (
	"url-shortener/handlers"
	"url-shortener/services"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func SetupRoutes(rdb *redis.Client) *gin.Engine {
	r := gin.Default()

	urlService := services.NewURLService(rdb)
	urlHandler := handlers.NewURLHandler(urlService)

	r.POST("/shorten", urlHandler.Shorten)
	r.GET("/:code", urlHandler.Redirect)

	return r
}
