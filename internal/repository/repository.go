package repository

import "github.com/nanmenkaimak/final-go-kbtu/internal/models"

type DatabaseRepo interface {
	InsertUser(newUser models.User) (int, error)
	GetUserByID(id int) (models.User, error)
	GetUsersByRole(role int) ([]models.User, error)
	UpdateUser(id int, updatedUser models.User) error
	DeleteUser(id int) error

	InsertItem(newItem models.Item) (int, error)
	GetAllItems() ([]models.Item, error)
	GetItemById(id int) (models.Item, error)
	GetItemsByName(name string) ([]models.Item, error)
	GetItemsByCategory(category string) ([]models.Item, error)
	GetItemsByPrice(price float64) ([]models.Item, error)
	GetItemsByRating(rating float64) ([]models.Item, error)
	UpdateItem(id int, updatedItem models.Item) error
	UpdateItemRating(id int, rating int) error
	DeleteItem(id int) error
	SortItemByPriceAsc() ([]models.Item, error)
	SortItemByPriceDesc() ([]models.Item, error)

	InsertComment(newComment models.Comment) (int, error)
	GetAllCommentsOfItem(itemID int) ([]models.Comment, error)

	Authenticate(email string, password string) (int, string, error)
}
