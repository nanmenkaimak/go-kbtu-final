package handlers

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func (m *Repository) DeleteSingleItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := m.DB.DeleteItem(id)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/items/filter", http.StatusSeeOther)
}
