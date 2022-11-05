package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vernandodev/pustaka-api/book"
	"github.com/vernandodev/pustaka-api/handler"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:dev034@@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("DB connection Success")

	db.AutoMigrate(&book.Book{})

	// Book Repository
	bookRepository := book.NewRepository(db)
	bookServices := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookServices)

	router := gin.Default()

	// API VERSIONING
	v1 := router.Group("/v1")
	v2 := router.Group("/v2")

	v1.GET("/", bookHandler.RootHandler)
	v2.GET("/hello", bookHandler.HelloHandler)
	router.GET("/books/:id/:title", bookHandler.BooksHandler)
	router.GET("/query", bookHandler.QueryHandler)
	router.POST("/books", bookHandler.PostBooksHandler)

	router.Run()
}

// LAYER
// Main
// Handler
// Service
// Repository
// DB
// MySQL
