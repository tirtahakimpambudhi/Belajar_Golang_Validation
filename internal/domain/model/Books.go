package model

import "time"

type Book struct {
	ISBN             string     `json:"isbn" validate:"required,isbn"`
	Title            string     `json:"title" validate:"required,min=1,max=255"`
	Author           string     `json:"author" validate:"required,min=1,max=255"`
	PublicationYears string     `json:"publication_years" validate:"required,min=4"`
	StatusBorrow     bool       `json:"status_borrow" validate:"boolean"`
	CreatedAt        *time.Time `json:"createdAt"`
	UpdatedAt        *time.Time `json:"updatedAt"`
}

func NewBook(ISBN string, title string, author string, publicationYears string, statusBorrow bool, createdAt *time.Time, updatedAt *time.Time) *Book {
	return &Book{ISBN: ISBN, Title: title, Author: author, PublicationYears: publicationYears, StatusBorrow: statusBorrow, CreatedAt: createdAt, UpdatedAt: updatedAt}
}

//Buat Custom validation default value
