package storage

import (
	"csuProject/models"
	"github.com/jmoiron/sqlx"
)

type LinkRepository struct {
	db *sqlx.DB
}

func NewLinkRepository(db *sqlx.DB) *LinkRepository {
	return &LinkRepository{db: db}
}

func (l LinkRepository) CreateLink(link string) (int, error) {
	var id int
	row := l.db.QueryRowx("INSERT INTO distr.Links (link) values ($1) returning id", link)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (l LinkRepository) GetLink(id int) (models.Link, error) {
	var link models.Link
	err := l.db.Get(&link, "select * from distr.Links where id=$1", id)
	if err != nil {
		return models.Link{}, err
	}
	return link, nil
}

func (l LinkRepository) UpdateLink(id int, status string) {
	l.db.QueryRow("UPDATE distr.Links SET status=$1 WHERE id=$2", status, id)
}
