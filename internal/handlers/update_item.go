package handlers

import (
	"github.com/nanmenkaimak/final-go-kbtu/internal/models"
	"github.com/nanmenkaimak/final-go-kbtu/internal/render"
	"log"
	"net/http"
	"path"
	"strconv"
)

func (m *Repository) ShowUpdateItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))

	item, err := m.DB.GetItemById(id)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	categories, err := m.DB.GetAllCategories()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["item"] = item
	data["categories"] = categories

	render.Template(w, r, "update.page.gohtml", &models.TemplateData{
		Form: nil,
		Data: data,
	})
}

func (m *Repository) UpdateItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
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
	categoryID, err := m.DB.GetIDOfCategoryByName(category)

	newItem := models.Items{
		Name:        name,
		Price:       price,
		CategoryID:  categoryID.ID,
		Description: description,
	}

	err = m.DB.UpdateItem(id, newItem)

	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/items/filter", http.StatusSeeOther)
}
