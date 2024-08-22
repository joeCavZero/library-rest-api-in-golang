package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/joeCavZero/library-rest-api-in-golang/internal/model"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/util/logkit"
)

var db *sql.DB

func init() {
	setupDatabase()
}

func setupDatabase() {
	lk := logkit.New("setup database")

	var err error
	db, err = sql.Open("sqlite3", "./db/library.db")
	if err != nil {
		lk.Error(err)
		panic(err)
	}

	if err = db.Ping(); err != nil {
		lk.Error(err)
		panic(err)
	}

	lk.Info("Database connected")

	_, err = db.Exec(
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

	lk.Info("Table books created successfully")

}

func CreateBook(book model.Book) (model.BookResponse, error) {
	lk := logkit.New("CreateBook")

	book_response := model.BookResponse{Book: book}

	result, err := db.Exec(
		`INSERT INTO books (title, author, location) VALUES (?,?,?)`,
		book.Title, book.Author, book.Location,
	)
	if err != nil {
		return book_response, err
	}

	lk.Info("Book inserted into table successfully")

	book_response.Id, err = result.LastInsertId()
	return book_response, nil
}

/*
 * The difference between Exec and Query is
 * that Exec is used for queries that do not return rows,
 * such as INSERT, UPDATE, DELETE, and CREATE TABLE.
 * Query is used for queries that return rows,
 */

func GetBooks() ([]model.Book, error) {
	lk := logkit.New("ReadBooks")

	rows, err := db.Query("SELECT title, author, location FROM books")
	if err != nil {
		lk.Error(err)
		return nil, err
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var new_book model.Book
		err = rows.Scan(&new_book.Title, &new_book.Author, &new_book.Location)
		if err != nil {
			lk.Error(err)
			return nil, err
		}
		books = append(books, new_book)
	}

	lk.Info("Books read from table successfully")
	return books, nil

}
