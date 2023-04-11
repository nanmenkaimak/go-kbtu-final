package handlers

import (
	"errors"
	"github.com/nanmenkaimak/final-go-kbtu/internal/models"
	"log"
	"net/http"
)

func (m *Repository) GetUserFromSession(w http.ResponseWriter, r *http.Request) (models.Users, error) {
	currentUserID, ok := m.App.Session.Get(r.Context(), "user_id").(int)
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return models.Users{}, errors.New("cannot find user in session")
	}

	currentUser, err := m.DB.GetUserByID(currentUserID)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return models.Users{}, err
	}
	return currentUser, nil
}
