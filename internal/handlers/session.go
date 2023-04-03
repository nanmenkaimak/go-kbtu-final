package handlers

import (
	"errors"
	"github.com/nanmenkaimak/final-go-kbtu/internal/models"
	"log"
	"net/http"
)

func (m *Repository) GetUserFromSession(w http.ResponseWriter, r *http.Request) (models.User, error) {
	currentUserID, ok := m.App.Session.Get(r.Context(), "user_id").(int)
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return models.User{}, errors.New("cannot find user in session")
	}

	currentUser, err := m.DB.GetUserByID(currentUserID)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return models.User{}, err
	}
	return currentUser, nil
}
