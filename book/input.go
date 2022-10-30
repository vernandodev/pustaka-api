package book

import "encoding/json"

type BookInput struct {
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"` // memakai json.Number akan mengkonversi input dari string menjadi number karena menggunakan validator.ValidatorErrors
	Email string      `json:"email" binding:"required,email"`
}
