package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey; autoIncrement:1"`
	FirstName string    `json:"first_name" gorm:"not null"`
	LastName  string    `json:"last_name" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Password  string    `json:"password" gorm:"not null"`
	Role      int       `json:"role" gorm:"not null"` // 1 - client, 2 - seller, 3 - admin
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

type Item struct {
	ID              int       `json:"id" gorm:"primaryKey; autoIncrement:1"`
	Name            string    `json:"name" gorm:"not null"`
	Price           float64   `json:"price" gorm:"not null"`
	Rating          float64   `json:"rating" gorm:"not null; default:0"`
	NumberOfRatings int       `json:"number_of_ratings" gorm:"not null; default:0"`
	Category        string    `json:"category" gorm:"not null"`
	Description     string    `json:"description" gorm:"not null"`
	SellerID        int       `json:"seller_id" gorm:"not null"`
	Seller          User      `json:"seller" gorm:"not null; constraint:OnDelete:CASCADE;"`
	Comments        []Comment `json:"comments" gorm:"constraint:OnDelete:CASCADE;"`
	CreatedAt       time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"not null"`
}

type Comment struct {
	ID         int       `json:"id" gorm:"primaryKey; autoIncrement:1"`
	ItemID     int       `json:"item_id" gorm:"not null"`
	Text       string    `json:"text" gorm:"not null"`
	Rating     int       `json:"rating" gorm:"not null"`
	AuthorID   int       `json:"author_id" gorm:"not null"`
	AuthorName string    `json:"author_name" gorm:"not null"`
	Author     User      `json:"author" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"not null"`
}
