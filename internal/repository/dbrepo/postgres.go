package dbrepo

import (
	"errors"
	"github.com/nanmenkaimak/final-go-kbtu/internal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (m *postgresDBRepo) InsertUser(newUser models.User) (int, error) {
	password, err := hashPassword(newUser.Password)
	if err != nil {
		return 0, err
	}
	newUser.Password = password

	result := m.DB.Create(&newUser)

	return newUser.ID, result.Error
}

func (m *postgresDBRepo) InsertItem(newItem models.Item) (int, error) {
	result := m.DB.Create(&newItem)

	return newItem.ID, result.Error
}

func (m *postgresDBRepo) InsertComment(newComment models.Comment) (int, error) {
	result := m.DB.Create(&newComment)

	return newComment.ID, result.Error
}

func (m *postgresDBRepo) GetAllItems() ([]models.Item, error) {
	var items []models.Item
	result := m.DB.Find(&items)

	if result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}

func (m *postgresDBRepo) GetUserByID(id int) (models.User, error) {
	var user models.User
	result := m.DB.Where("id = ?", id).Find(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (m *postgresDBRepo) GetUsersByRole(role int) ([]models.User, error) {
	var user []models.User
	result := m.DB.Where("role = ?", role).Find(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (m *postgresDBRepo) GetItemById(id int) (models.Item, error) {
	var item models.Item
	result := m.DB.Where("id = ?", id).Find(&item)

	if result.Error != nil {
		return item, result.Error
	}

	return item, nil
}

func (m *postgresDBRepo) GetItemsByName(name string) ([]models.Item, error) {
	var item []models.Item
	result := m.DB.Where("name = ?", name).Find(&item)

	if result.Error != nil {
		return item, result.Error
	}

	return item, nil
}

func (m *postgresDBRepo) GetItemsByCategory(category string) ([]models.Item, error) {
	var item []models.Item
	result := m.DB.Where("category = ?", category).Find(&item)

	if result.Error != nil {
		return item, result.Error
	}

	return item, nil
}

func (m *postgresDBRepo) GetItemsByPrice(price float64) ([]models.Item, error) {
	var item []models.Item
	result := m.DB.Where("price <= ?", price).Find(&item)

	if result.Error != nil {
		return item, result.Error
	}

	return item, nil
}

func (m *postgresDBRepo) GetItemsByRating(rating float64) ([]models.Item, error) {
	var item []models.Item
	result := m.DB.Where("rating >= ?", rating).Find(&item)

	if result.Error != nil {
		return item, result.Error
	}

	return item, nil
}

func (m *postgresDBRepo) GetAllCommentsOfItem(itemID int) ([]models.Comment, error) {
	var comments []models.Comment
	result := m.DB.Where("item_id = ?", itemID).Find(&comments)

	if result.Error != nil {
		return comments, result.Error
	}

	return comments, nil
}

//func (m *postgresDBRepo) GetAllCommentsOfUser(userID int) ([]models.Comment, error) {
//
//}
//
//func (m *postgresDBRepo) GetAllCommentsByRating(rating float64) ([]models.Comment, error) {
//
//}

func (m *postgresDBRepo) UpdateUser(id int, updatedUser models.User) error {
	var newUser models.User
	password, err := hashPassword(updatedUser.Password)
	if err != nil {
		return err
	}
	updatedUser.Password = password
	result := m.DB.Model(&newUser).Where("id = ?", id).Updates(updatedUser)

	return result.Error
}

func (m *postgresDBRepo) UpdateItem(id int, updatedItem models.Item) error {
	var newItem models.Item

	result := m.DB.Model(&newItem).Where("id = ?", id).Updates(updatedItem)

	return result.Error
}

func (m *postgresDBRepo) UpdateItemRating(id int, rating int) error {
	var item models.Item

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
	var user models.User
	result := m.DB.Delete(&user, id)

	return result.Error
}

func (m *postgresDBRepo) DeleteItem(id int) error {
	var item models.Item

	result := m.DB.Delete(&item, id)

	return result.Error
}

//func (m *postgresDBRepo) DeleteComment(id int) error {
//
//}

func (m *postgresDBRepo) SortItemByPriceAsc() ([]models.Item, error) {
	var items []models.Item

	result := m.DB.Order("price asc").Find(&items)

	return items, result.Error
}

func (m *postgresDBRepo) SortItemByPriceDesc() ([]models.Item, error) {
	var items []models.Item

	result := m.DB.Order("price desc").Find(&items)

	return items, result.Error
}

func (m *postgresDBRepo) Authenticate(email string, password string) (int, string, error) {
	var user models.User

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
