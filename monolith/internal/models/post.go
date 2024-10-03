package models

import "time"

//go:generate easytags $GOFILE json,mapper
type Post struct {
	ID               int       `json:"id" mapper:"id"`
	Title            string    `json:"title" mapper:"title"`
	ShortDescription string    `json:"short_description" mapper:"short_description"`
	FullDescription  string    `json:"full_description" mapper:"full_description"`
	Author           int       `json:"author" mapper:"author"`
	CreatedAt        time.Time `json:"created_at" mapper:"created_at"`
	UpdateAt         time.Time `json:"updateAt " mapper:"update_at"`
	DeletedAt        time.Time `json:"delete_st" mapper:"delete_st"`
}
