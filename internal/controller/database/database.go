package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/joeCavZero/library-rest-api-in-golang/config"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/model"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/util/logkit"
)

type databaseController struct {
	db *sql.DB
}

func NewDatabaseController() *databaseController {
	lk := logkit.New("NEW DATABASE CONTROLLER")
	var err error

	controller := databaseController{}
	controller.db, err = sql.Open("sqlite3", config.DB_PATH)
	if err != nil {
		lk.Error(err)
		panic(err)
	}

	if err = controller.db.Ping(); err != nil {
		lk.Error(err)
		panic(err)
	}

	lk.Info("Database connected successfully")

	_, err = controller.db.Exec(
		`CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title VARCHAR(50) NOT NULL,
			author VARCHAR(50) NOT NULL,
			location VARCHAR(50) NOT NULL
		)`,
	)

	if err != nil {
		lk.Error(err)
		panic(err)
	}

	return &controller

}

func (d *databaseController) CreateBook(book model.Book) (model.BookResponse, error) {
	lk := logkit.New("CreateBook")

	book_response := model.BookResponse{Book: book}

	result, err := d.db.Exec(
		`INSERT INTO books (title, author, location) VALUES (?,?,?)`,
		book.Title, book.Author, book.Location,
	)
	if err != nil {
		return model.BookResponse{}, err
	}

	lk.Info("Book inserted into table successfully")

	last_ins_id, err := result.LastInsertId()
	if err != nil {
		return model.BookResponse{}, err
	}

	book_response.Id = uint32(last_ins_id)

	return book_response, nil
}

/*
 * The difference between Exec and Query is
 * that Exec is used for queries that do not return rows,
 * such as INSERT, UPDATE, DELETE, and CREATE TABLE.
 * Query is used for queries that return rows,
 */

func (d *databaseController) GetBooks() ([]model.BookResponse, error) {
	lk := logkit.New("ReadBooks")

	rows, err := d.db.Query("SELECT id, title, author, location FROM books")
	if err != nil {
		lk.Error(err)
		return nil, err
	}
	defer rows.Close()

	var books []model.BookResponse
	for rows.Next() {
		var new_book model.BookResponse
		err = rows.Scan(&new_book.Id, &new_book.Title, &new_book.Author, &new_book.Location)
		if err != nil {
			lk.Error(err)
			return nil, err
		}
		books = append(books, new_book)
	}

	lk.Info("Books read from table successfully")
	return books, nil

}

func (d *databaseController) GetBookById(id uint32) (model.Book, error) {
	lk := logkit.New("Get-Book-By-Id")
	lk.Infof("id: %d", id)
	var new_book model.Book
	row := d.db.QueryRow("SELECT title, author, location FROM books WHERE id = ?", id)

	lk.Infof("row: %v", row)

	err := row.Scan(&new_book.Title, &new_book.Author, &new_book.Location)
	if err != nil {
		return model.Book{}, err
	}

	return new_book, nil
}

func (d *databaseController) DeleteBookById(id uint32) error {
	_, err := d.db.Exec("DELETE FROM books WHERE id = ?", id)
	return err
}

func (d *databaseController) UpdateBookById(id uint32, book model.Book) (model.BookResponse, error) {
	query := "UPDATE books SET title = ?, author = ?, location = ? WHERE id = ?"
	_, err := d.db.Exec(query, book.Title, book.Author, book.Location, id)
	if err != nil {
		return model.BookResponse{}, err
	}

	new_book, err := d.GetBookById(id)
	if err != nil {
		return model.BookResponse{}, err
	}

	response := model.BookResponse{
		Id:   id,
		Book: new_book,
	}

	return response, nil
}
