package repository

import (
	"github.com/setiawannuha/go_book/internal/models"

	"github.com/jmoiron/sqlx"
)

type BookRepositoryInterface interface {
	CreateData(body *models.Book) (string, error)
	GetAllData() (*models.Books, error)
	GetDetailData(id string) (*models.Book, error)
	UpdateData(id string, body *models.Book) (string, error)
	DeleteData(id string) (string, error)
}

type BookRepository struct {
	*sqlx.DB
}

func NewBookRepository(db *sqlx.DB) *BookRepository {
	return &BookRepository{db}
}

func (r *BookRepository) CreateData(body *models.Book) (string, error) {
	query := `
		INSERT INTO tbl_books(
			title, image, author, release_date)
			VALUES (:title, :image, :author, :release_date)
	`
	_, err := r.NamedExec(query, body)
	if err != nil {
		return "", err
	}
	return "data created", nil
}

func (r *BookRepository) GetAllData() (*models.Books, error) {
	query := `SELECT * FROM tbl_books`
	data := models.Books{}
	err := r.Select(&data, query)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *BookRepository) GetDetailData(id string) (*models.Book, error) {
	query := `SELECT * FROM tbl_books WHERE id=$1`
	data := models.Book{}
	err := r.Get(&data, query, id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *BookRepository) UpdateData(id string, body *models.Book) (string, error) {
	query := `UPDATE tbl_books SET title = :title, image = :image, author = :author WHERE id = :id`
	params := map[string]interface{}{
		"id":     id,
		"title":  body.Title,
		"image":  body.Image,
		"author": body.Author,
	}
	_, err := r.NamedExec(query, params)
	if err != nil {
		return "", err
	}
	return "data updated", nil
}

func (r *BookRepository) DeleteData(id string) (string, error) {
	query := `DELETE FROM tbl_books WHERE id=$1`
	_, err := r.Exec(query, id)
	if err != nil {
		return "", err
	}
	return "data deleted", nil

}
