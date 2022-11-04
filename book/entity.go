package book

import "time"

type Book struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Price       int
	Email       string
	Description string
	Rating      int
	Discount    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
