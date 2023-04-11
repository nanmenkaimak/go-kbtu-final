package dbrepo

import (
	"errors"
	"github.com/nanmenkaimak/final-go-kbtu/internal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (m *postgresDBRepo) InsertUser(newUser models.Users) (int, error) {
	password, err := hashPassword(newUser.Password)
	if err != nil {
		return 0, err
	}
	newUser.Password = password

	result := m.DB.Create(&newUser)

	return newUser.ID, result.Error
}

func (m *postgresDBRepo) InsertItem(newItem models.Items) (int, error) {
	result := m.DB.Create(&newItem)

	return newItem.ID, result.Error
}

func (m *postgresDBRepo) InsertComment(newComment models.Comments) (int, error) {
	result := m.DB.Create(&newComment)

	return newComment.ID, result.Error
}

func (m *postgresDBRepo) GetAllItems() ([]models.Items, error) {
	var items []models.Items
	result := m.DB.Find(&items)

	for i := range items {
		cat, _ := m.GetNameOfCategoryByID(items[i].CategoryID)
		items[i].Category = cat
	}

	return items, result.Error
}

func (m *postgresDBRepo) GetUserByID(id int) (models.Users, error) {
	var user models.Users
	result := m.DB.Where("id = ?", id).Find(&user)

	return user, result.Error
}

func (m *postgresDBRepo) GetUsersByRole(role int) ([]models.Users, error) {
	var user []models.Users
	result := m.DB.Where("role_id = ?", role).Find(&user)

	return user, result.Error
}

func (m *postgresDBRepo) GetItemById(id int) (models.Items, error) {
	var item models.Items
	result := m.DB.Where("id = ?", id).Find(&item)

	cat, _ := m.GetNameOfCategoryByID(item.CategoryID)
	item.Category = cat

	return item, result.Error
}

func (m *postgresDBRepo) GetItemsByName(name string) ([]models.Items, error) {
	var item []models.Items
	result := m.DB.Where("name = ?", name).Find(&item)

	return item, result.Error
}

func (m *postgresDBRepo) GetIDOfCategoryByName(name string) (models.Categories, error) {
	var category models.Categories
	result := m.DB.Where("name = ?", name).Find(&category)

	return category, result.Error
}

func (m *postgresDBRepo) GetNameOfCategoryByID(id int) (models.Categories, error) {
	var category models.Categories
	result := m.DB.Where("id = ?", id).Find(&category)

	return category, result.Error
}

func (m *postgresDBRepo) GetIDOfRoleByName(name string) (models.Roles, error) {
	var role models.Roles
	result := m.DB.Where("name = ?", name).Find(&role)

	return role, result.Error
}

func (m *postgresDBRepo) GetAllCategories() ([]models.Categories, error) {
	var categories []models.Categories
	result := m.DB.Find(&categories)

	return categories, result.Error
}

func (m *postgresDBRepo) GetItemsByCategory(name string) ([]models.Items, error) {
	category, err := m.GetIDOfCategoryByName(name)
	if err != nil {
		return nil, err
	}
	var item []models.Items
	result := m.DB.Where("category_id = ?", category.ID).Find(&item)

	return item, result.Error
}

func (m *postgresDBRepo) GetItemsByPrice(price float64) ([]models.Items, error) {
	var item []models.Items
	result := m.DB.Where("price <= ?", price).Find(&item)

	return item, result.Error
}

func (m *postgresDBRepo) GetItemsByRating(rating float64) ([]models.Items, error) {
	var item []models.Items
	result := m.DB.Where("rating >= ?", rating).Find(&item)

	return item, result.Error
}

func (m *postgresDBRepo) GetAllCommentsOfItem(itemID int) ([]models.Comments, error) {
	var comments []models.Comments
	result := m.DB.Where("item_id = ?", itemID).Find(&comments)

	return comments, result.Error
}

//func (m *postgresDBRepo) GetAllCommentsOfUser(userID int) ([]models.Comment, error) {
//
//}
//
//func (m *postgresDBRepo) GetAllCommentsByRating(rating float64) ([]models.Comment, error) {
//
//}

func (m *postgresDBRepo) UpdateUser(id int, updatedUser models.Users) error {
	var newUser models.Users
	password, err := hashPassword(updatedUser.Password)
	if err != nil {
		return err
	}
	updatedUser.Password = password
	result := m.DB.Model(&newUser).Where("id = ?", id).Updates(updatedUser)

	return result.Error
}

func (m *postgresDBRepo) UpdateItem(id int, updatedItem models.Items) error {
	var newItem models.Items

	result := m.DB.Model(&newItem).Where("id = ?", id).Updates(updatedItem)

	return result.Error
}

func (m *postgresDBRepo) UpdateItemRating(id int, rating int) error {
	var item models.Items

	query := `update items set rating = rating + ((? - rating) / (number_of_ratings + 1)),
			number_of_ratings = number_of_ratings + 1,
			updated_at = ?
			where id = ?`

	result := m.DB.Raw(query, rating, time.Now(), id).Scan(&item)

	return result.Error
}

//func (m *postgresDBRepo) UpdateComment(updatedItem models.Comment) error {
//
//}
//

func (m *postgresDBRepo) DeleteUser(id int) error {
	var user models.Users
	result := m.DB.Delete(&user, id)

	return result.Error
}

func (m *postgresDBRepo) DeleteItem(id int) error {
	var item models.Items

	result := m.DB.Delete(&item, id)

	return result.Error
}

//func (m *postgresDBRepo) DeleteComment(id int) error {
//
//}

func (m *postgresDBRepo) SortItemByPriceAsc() ([]models.Items, error) {
	var items []models.Items

	result := m.DB.Order("price asc").Find(&items)

	return items, result.Error
}

func (m *postgresDBRepo) SortItemByPriceDesc() ([]models.Items, error) {
	var items []models.Items

	result := m.DB.Order("price desc").Find(&items)

	return items, result.Error
}

func (m *postgresDBRepo) Authenticate(email string, password string) (int, string, error) {
	var user models.Users

	m.DB.First(&user, "email = ?", email)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("incorrect password")
	} else if err != nil {
		return 0, "", err
	}
	return user.ID, user.Password, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
