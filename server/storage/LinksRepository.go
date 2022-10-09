package storage

import "github.com/jmoiron/sqlx"

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

func (l LinkRepository) GetLink(id int) (string, error) {
	var link string
	err := l.db.Get(&link, "select link from distr.Links where id=$1", id)
	if err != nil {
		return "", err
	}
	return link, nil
}
