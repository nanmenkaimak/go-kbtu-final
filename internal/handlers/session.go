package handlers

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nanmenkaimak/final-go-kbtu/internal/models"
	"net/http"
)

func (m *Repository) GetUserFromJWT(w http.ResponseWriter, r *http.Request) (models.Users, error) {
	c, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return models.Users{}, errors.New("cannot find user in session")
	}

	tknStr := c.Value

	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !tkn.Valid {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return models.Users{}, errors.New("cannot find user in session")
	}

	user, err := m.DB.GetUserByID(claims.ID)

	return user, err
}
