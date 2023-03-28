package dbrepo

import (
	"github.com/nanmenkaimak/final-go-kbtu/internal/config"
	"github.com/nanmenkaimak/final-go-kbtu/internal/repository"
	"gorm.io/gorm"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *gorm.DB
}

func NewPostgresRepo(conn *gorm.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
