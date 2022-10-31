package book

import "time"

type Book struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Price       int
	Rating      int
	Discount    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
