package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jojonicho/s16-movie/api"
	"github.com/jojonicho/s16-movie/movie"
)

func InitServer(apiKey string) *gin.Engine {
	movie := movie.NewService(apiKey)
	movieHandler := api.NewMovieHTTPHandler(movie)

	router := gin.Default()

	movieGroup := router.Group("/search")
	movieGroup.GET("/", movieHandler.Search)

	detailGroup := router.Group("/detail")
	detailGroup.GET("/", movieHandler.Detail)

	return router
}

func main() {
	apiKey := os.Getenv("OMDB_KEY")
	if apiKey == "" {
		panic("OMDB_KEY env variable not set")
	}
	svr := InitServer(apiKey)

	if err := svr.Run(); err != nil {
		log.Fatalf("Failed starting http server: %v", err)
	}
}
