package handlers

import (
	"net/http"
	"time"
)

func (m *Repository) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
