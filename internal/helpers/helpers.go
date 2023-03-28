package helpers

import (
	"github.com/nanmenkaimak/final-go-kbtu/internal/config"
	"net/http"
)

var app *config.AppConfig

func IsAuthenticated(r *http.Request) bool {
	exists := app.Session.Exists(r.Context(), "user_id")
	return exists
}
