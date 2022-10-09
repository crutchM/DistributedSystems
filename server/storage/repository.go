package storage

import "github.com/jmoiron/sqlx"

type LinksRepo interface {
	CreateLink(link string) (int, error)
	GetLink(id int) (string, error)
}

type Repository struct {
	LinksRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{LinksRepo: NewLinkRepository(db)}
}
