package models

import "time"

type Book struct {
	Id           string     `db:"id" json:"id" form:"id" valid:"-"`
	Title        string     `db:"title" json:"title" form:"title" valid:"stringlength(5|100)~Title minimal 5 dan maksimal 100"`
	Image        string     `db:"image" json:"image" valid:"-"`
	Author       string     `db:"author" json:"author" form:"author" valid:"type(string)"`
	Release_date string     `db:"release_date" json:"release_date" form:"release_date" valid:"-"`
	Created_at   *time.Time `db:"created_at" json:"created_at" valid:"-"`
	Updated_at   *time.Time `db:"updated_at" json:"updated_at" valid:"-"`
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
