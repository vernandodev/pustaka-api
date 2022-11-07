package book

import "encoding/json"

type BookRequest struct {
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"` // memakai json.Number akan mengkonversi input dari string menjadi number karena menggunakan validator.ValidatorErrors
	Email       string      `json:"email" binding:"required,email"`
	Description string      `json:"description" binding:"required"`
	Rating      json.Number `json:"rating" binding:"required,number"`
	Discount    json.Number `json:"discount" binding:"required,number"`
}
