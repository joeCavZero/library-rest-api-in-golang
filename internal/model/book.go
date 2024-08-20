package model

type BookResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Location string `json:"location"`
}
