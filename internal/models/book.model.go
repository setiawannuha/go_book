package models

import "time"

type Book struct {
	Id           string     `db:"id" json:"id"`
	Title        string     `db:"title" json:"title"`
	Image        string     `db:"image" json:"image"`
	Author       string     `db:"author" json:"author"`
	Release_date string     `db:"release_date" json:"release_date"`
	Created_at   *time.Time `db:"created_at" json:"created_at"`
	Updated_at   *time.Time `db:"updated_at" json:"updated_at"`
}

type Books []Book
