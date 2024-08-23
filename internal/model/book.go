package model

import "errors"

type Book struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Location string `json:"location"`
}

func (b Book) Validate() error {
	if b.Title == "" || b.Author == "" || b.Location == "" {
		return errors.New("invalid book camp fields")
	}
	return nil
}

type BookResponse struct {
	Id uint32 `json:"id"`
	Book
}
