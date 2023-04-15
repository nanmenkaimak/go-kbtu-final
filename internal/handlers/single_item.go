package handlers

import (
	"github.com/nanmenkaimak/final-go-kbtu/internal/models"
	"github.com/nanmenkaimak/final-go-kbtu/internal/render"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"
)

func (m *Repository) SingleItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))
	item, err := m.DB.GetItemById(id)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	comments, err := m.DB.GetAllCommentsOfItem(id)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	for i := range comments {
		author, _ := m.DB.GetUserByID(comments[i].AuthorID)
		comments[i].Author.FirstName = author.FirstName
	}

	currentUser, _ := m.GetUserFromJWT(w, r)

	data := make(map[string]interface{})
	data["item"] = item
	data["comments"] = comments
	data["user"] = currentUser
	data["theircomment"] = false

	if item.SellerID == currentUser.ID || currentUser.RoleID == 3 {
		data["theircomment"] = true
	}

	render.Template(w, r, "item.page.gohtml", &models.TemplateData{
		Form: nil,
		Data: data,
	})
}

func (m *Repository) PostSingleItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	rate, err := strconv.Atoi(r.Form.Get("rating"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	comment := r.Form.Get("comment")

	if comment == "" {
		log.Println(err)
		http.Redirect(w, r, "/items/filter", http.StatusSeeOther)
		return
	}

	err = m.DB.UpdateItemRating(id, rate)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	currentUser, _ := m.GetUserFromJWT(w, r)

	newComment := models.Comments{
		ItemID:    id,
		Text:      comment,
		Rating:    rate,
		AuthorID:  currentUser.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = m.DB.InsertComment(newComment)

	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/items/filter", http.StatusSeeOther)
}
