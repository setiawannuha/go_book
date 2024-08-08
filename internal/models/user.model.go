package models

import "time"

type User struct {
	Id         string     `db:"id" json:"id" valid:"-"`
	Email      string     `db:"email" json:"email" valid:"email"`
	Password   string     `db:"password" json:"password" valid:"stringlength(5|256)~Password minimal 5 karakter"`
	Created_at *time.Time `db:"created_at" json:"created_at" valid:"-"`
	Updated_at *time.Time `db:"updated_at" json:"updated_at" valid:"-"`
}

type Users []User

// CREATE TABLE public.tbl_users
// (
//     id character varying COLLATE pg_catalog."default" NOT NULL DEFAULT uuid_generate_v4(),
//     email character varying COLLATE pg_catalog."default" NOT NULL,
//     password character varying COLLATE pg_catalog."default" NOT NULL,
//     created_at timestamp with time zone DEFAULT now(),
//     updated_at timestamp with time zone DEFAULT now(),
//     CONSTRAINT tbl_users_pkey PRIMARY KEY (id)
// )
