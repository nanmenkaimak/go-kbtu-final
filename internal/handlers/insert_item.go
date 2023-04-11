package handlers

import (
	"github.com/nanmenkaimak/final-go-kbtu/internal/forms"
	"github.com/nanmenkaimak/final-go-kbtu/internal/models"
	"github.com/nanmenkaimak/final-go-kbtu/internal/render"
	"log"
	"net/http"
	"strconv"
)

func (m *Repository) ShowInsertItem(w http.ResponseWriter, r *http.Request) {
	categories, err := m.DB.GetAllCategories()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["categories"] = categories
	render.Template(w, r, "insert.page.gohtml", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

func (m *Repository) InsertItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	name := r.Form.Get("name")
	price, err := strconv.ParseFloat(r.Form.Get("price"), 64)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	category := r.Form.Get("category")
	description := r.Form.Get("description")

	currentUser, _ := m.GetUserFromSession(w, r)

	categoryID, err := m.DB.GetIDOfCategoryByName(category)

	newItem := models.Items{
		Name:        name,
		Price:       price,
		CategoryID:  categoryID.ID,
		Description: description,
		SellerID:    currentUser.ID,
	}

	_, err = m.DB.InsertItem(newItem)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/items/filter", http.StatusSeeOther)
}
