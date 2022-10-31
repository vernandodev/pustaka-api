package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vernandodev/pustaka-api/book"
	"github.com/vernandodev/pustaka-api/controllers"
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

	// CRUD

	// CREATE
	// membuat object baru dari struct
	// book := book.Book{}
	// book.Title = "Atomic Habbit"
	// book.Description = "Buku cerita 2"
	// book.Price = 125000
	// book.Rating = 9
	// book.Discount = 5

	// err = db.Create(&book).Error
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Success Create")

	// READ
	// var books []book.Book
	// err = db.Debug().Where("rating = ?", 5).Find(&books).Error
	// if err != nil {
	// 	panic(err)
	// }
	// for _, b := range books {
	// 	fmt.Println(b.Title)
	// 	fmt.Println(b)
	// 	fmt.Println("################")
	// }

	// UPDATE
	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	panic(err)
	// }

	// book.Title = "Man Tiger (Revised Edition)"
	// db.Save(&book)
	// if err != nil {
	// 	panic(err)
	// }

	// DELETE
	var book book.Book

	err = db.Debug().Where("id = ?", 3).First(&book).Error
	if err != nil {
		panic(err)
	}

	err = db.Debug().Delete(&book).Error
	if err != nil {
		panic(err)
	}

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
