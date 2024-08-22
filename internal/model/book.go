package model

import "errors"

type Book struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Location string `json:"location"`
}

func (b Book) Validate() error {
	if b.Title == "" || b.Author == "" || b.Location == "" {
		return errors.New("Invalid book camp fiels")
	}
	return nil
}

type BookResponse struct {
	Id int64 `json:"id"`
	Book
}
