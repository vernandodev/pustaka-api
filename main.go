package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vernandodev/pustaka-api/controllers"
)

func main() {
	router := gin.Default()

	// API VERSIONING

	v1 := router.Group("/v1")
	v2 := router.Group("/v2")

	v1.GET("/", controllers.RootHandler)
	v2.GET("/hello", controllers.HelloHandler)
	router.GET("/books/:id/:title", controllers.BooksHandler)
	router.GET("/query", controllers.QueryHandler)
	router.POST("/books", controllers.PostBooksHandler)

	router.Run()
}
