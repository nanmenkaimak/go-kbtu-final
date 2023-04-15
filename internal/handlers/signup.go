package handlers

import (
	"github.com/nanmenkaimak/final-go-kbtu/internal/forms"
	"github.com/nanmenkaimak/final-go-kbtu/internal/models"
	"github.com/nanmenkaimak/final-go-kbtu/internal/render"
	"net/http"
)

func (m *Repository) ShowSignUp(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "signup.page.gohtml", &models.TemplateData{
		Form: forms.New(nil),
	})
}

func (m *Repository) SignUp(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	first_name := r.Form.Get("first_name")
	last_name := r.Form.Get("last_name")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	role := r.Form.Get("role")

	roleID, err := m.DB.GetIDOfRoleByName(role)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	newUser := models.Users{
		FirstName: first_name,
		LastName:  last_name,
		Email:     email,
		Password:  password,
		RoleID:    roleID.ID,
	}

	form := forms.New(r.PostForm)

	form.IsEmail("email")
	form.MinLength("password", 8)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["user"] = newUser
		render.Template(w, r, "signup.page.gohtml", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	_, err = m.DB.InsertUser(newUser)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
