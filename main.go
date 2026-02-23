package main

import (
	"url-shortener/config"
	"url-shortener/routes"
)

func main() {
	rdb := config.ConnectRedis()
	r := routes.SetupRoutes(rdb)
	r.Run(":8080")
}
