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
