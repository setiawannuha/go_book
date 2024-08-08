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

// CREATE TABLE public.tbl_books
// (
//     id character varying COLLATE pg_catalog."default" NOT NULL DEFAULT uuid_generate_v4(),
//     title character varying COLLATE pg_catalog."default" NOT NULL,
//     author character varying COLLATE pg_catalog."default" NOT NULL,
//     release_date date NOT NULL,
//     created_at timestamp with time zone DEFAULT now(),
//     updated_at timestamp with time zone DEFAULT now(),
//     image character varying COLLATE pg_catalog."default",
//     CONSTRAINT tbl_books_pkey PRIMARY KEY (id)
// )
