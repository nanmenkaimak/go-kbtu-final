package models

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Roles struct {
	ID   int    `json:"id" gorm:"primaryKey; not null"`
	Name string `json:"name" gorm:"not null"`
}

type Users struct {
	ID        int       `json:"id" gorm:"primaryKey; autoIncrement:1"`
	FirstName string    `json:"first_name" gorm:"not null"`
	LastName  string    `json:"last_name" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Password  string    `json:"password" gorm:"not null"`
	RoleID    int       `json:"role_id" gorm:"not null"` // fk Roles id
	Role      Roles     `json:"role" gorm:"not null; constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

type Categories struct {
	ID   int    `json:"id" gorm:"primaryKey; not null"`
	Name string `json:"name" gorm:"not null"`
}

type Items struct {
	ID              int        `json:"id" gorm:"primaryKey; autoIncrement:1"`
	Name            string     `json:"name" gorm:"not null"`
	Description     string     `json:"description" gorm:"not null"`
	Price           float64    `json:"price" gorm:"not null"`
	SellerID        int        `json:"seller_id" gorm:"not null"` // fk User id
	Seller          Users      `json:"seller" gorm:"not null; constraint:OnDelete:CASCADE;"`
	CategoryID      int        `json:"category_id" gorm:"not null"` // fk Categories id
	Category        Categories `json:"category" gorm:"not null; constraint:OnDelete:CASCADE;"`
	Rating          float64    `json:"rating" gorm:"not null; default:0"`
	NumberOfRatings int        `json:"number_of_ratings" gorm:"not null; default:0"`
	CreatedAt       time.Time  `json:"created_at" gorm:"not null"`
	UpdatedAt       time.Time  `json:"updated_at" gorm:"not null"`
}

type Comments struct {
	ID        int       `json:"id" gorm:"primaryKey; autoIncrement:1"`
	ItemID    int       `json:"item_id" gorm:"not null"` // fk Item id
	Item      Items     `json:"item" gorm:"not null; constraint:OnDelete:CASCADE;"`
	Text      string    `json:"text" gorm:"not null"`
	Rating    int       `json:"rating" gorm:"not null"`
	AuthorID  int       `json:"author_id" gorm:"not null"` // fk User id
	Author    Users     `json:"author" gorm:"not null; constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

type Claims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}
