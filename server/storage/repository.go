package storage

import (
	"csuProject/server/models"
	"github.com/jmoiron/sqlx"
)

type LinksRepo interface {
	CreateLink(link string) (int, error)
	GetLink(id int) (models.Link, error)
	UpdateLink(id int, status string)
}

type Repository struct {
	LinksRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{LinksRepo: NewLinkRepository(db)}
}
