package handlers

import (
	"github.com/nanmenkaimak/final-go-kbtu/internal/models"
	"github.com/nanmenkaimak/final-go-kbtu/internal/render"
	"log"
	"net/http"
	"strconv"
)

func (m *Repository) GetAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := m.DB.GetAllItems()
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	currentUser, _ := m.GetUserFromSession(w, r)

	data := make(map[string]interface{})
	data["items"] = items
	data["user"] = currentUser

	render.Template(w, r, "filter.page.gohtml", &models.TemplateData{
		Form: nil,
		Data: data,
	})
}

// SortItems filters by name, rating, price
func (m *Repository) SortItems(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var items []models.Items

	if r.Form.Get("price") != "" {
		price, err := strconv.ParseFloat(r.Form.Get("price"), 64)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		items, err = m.DB.GetItemsByPrice(price)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	} else if r.Form.Get("rating") != "" {
		rating, err := strconv.ParseFloat(r.Form.Get("rating"), 64)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		items, err = m.DB.GetItemsByRating(rating)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	} else if r.Form.Get("name") != "" {
		name := r.Form.Get("name")

		items, err = m.DB.GetItemsByName(name)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	} else if r.Form.Get("sorting") != "" {
		sorting := r.Form.Get("sorting")

		if sorting == "desc" {
			items, err = m.DB.SortItemByPriceDesc()
		} else {
			items, err = m.DB.SortItemByPriceAsc()
		}
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	currentUser, _ := m.GetUserFromSession(w, r)

	for i := range items {
		cat, _ := m.DB.GetNameOfCategoryByID(items[i].CategoryID)
		items[i].Category = cat
	}

	data := make(map[string]interface{})
	data["items"] = items
	data["user"] = currentUser

	render.Template(w, r, "filter.page.gohtml", &models.TemplateData{
		Data: data,
	})
}
