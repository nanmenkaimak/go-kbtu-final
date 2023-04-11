package repository

import "github.com/nanmenkaimak/final-go-kbtu/internal/models"

type DatabaseRepo interface {
	InsertUser(newUser models.Users) (int, error)
	GetUserByID(id int) (models.Users, error)
	GetUsersByRole(role int) ([]models.Users, error)
	UpdateUser(id int, updatedUser models.Users) error
	DeleteUser(id int) error

	GetIDOfRoleByName(name string) (models.Roles, error)
	GetAllCategories() ([]models.Categories, error)
	GetNameOfCategoryByID(id int) (models.Categories, error)

	InsertItem(newItem models.Items) (int, error)
	GetAllItems() ([]models.Items, error)
	GetItemById(id int) (models.Items, error)
	GetItemsByName(name string) ([]models.Items, error)
	GetIDOfCategoryByName(name string) (models.Categories, error)
	GetItemsByCategory(category string) ([]models.Items, error)
	GetItemsByPrice(price float64) ([]models.Items, error)
	GetItemsByRating(rating float64) ([]models.Items, error)
	UpdateItem(id int, updatedItem models.Items) error
	UpdateItemRating(id int, rating int) error
	DeleteItem(id int) error
	SortItemByPriceAsc() ([]models.Items, error)
	SortItemByPriceDesc() ([]models.Items, error)

	InsertComment(newComment models.Comments) (int, error)
	GetAllCommentsOfItem(itemID int) ([]models.Comments, error)

	Authenticate(email string, password string) (int, string, error)
}
