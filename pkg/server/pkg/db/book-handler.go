package db

import (
	"book-management/pkg/apis"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Handler wraps the standard operation of the REST application for interacting with the database
type Handler interface {
	CreateBook(book *apis.Book) (message string, err error)
	UpdateBook(book *apis.Book) (message string, err error)
	GetBook(filters *apis.FilterChain) (books []apis.Book, err error)
}

// MySQLHandler is the wrapper for the MySQL database
type MySQLHandler struct {
	db *sql.DB
}

// NewMySQLHandler returns a new MySQLHandler and set up the connection to the database
func NewMySQLHandler(opts Options) (*MySQLHandler, error) {
	handler := &MySQLHandler{}

	config := mysql.NewConfig()
	config.Addr = opts.Address()
	config.User = opts.User
	config.Passwd = opts.Pass
	config.DBName = opts.DB
	config.ParseTime = true
	config.Loc = time.UTC

	connector, err := mysql.NewConnector(config)
	if err != nil {
		return nil, fmt.Errorf("set up db connection: %v", err)
	}

	handler.db = sql.OpenDB(connector)

	return handler, nil
}

// CreateBook creates a new book in the database
func (s *MySQLHandler) CreateBook(book *apis.Book) (message string, err error) {
	stmt, err := s.db.Prepare("INSERT INTO books (title, author, description, isbn, published_date, edition, genre) VALUES (?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		fmt.Println(fmt.Errorf("prepare statement: %v", err))
		return "", fmt.Errorf("internal error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(book.Title, book.Author, book.Description, book.Isbn, book.PublishedDate.String(), int64(book.Edition), book.Genre)
	if err != nil {
		fmt.Println(fmt.Errorf("execute statement: %v", err))
		return "", fmt.Errorf("internal error")
	}

	return fmt.Sprintf("Created book %v written by %v with ISBN: %v", book.Title, book.Author, book.Isbn), nil
}

// UpdateBook updates an existing book in the database
func (s *MySQLHandler) UpdateBook(book *apis.Book) (message string, err error) {
	stmt, err := s.db.Prepare("UPDATE books SET title = ?, author = ?, description = ?, published_date = ?, edition = ?, genre = ? WHERE isbn = ?")

	if err != nil {
		fmt.Println(fmt.Errorf("prepare statement: %v", err))
		return "", fmt.Errorf("internal error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(book.Title, book.Author, book.Description, book.PublishedDate.String(), int64(book.Edition), book.Genre, book.Isbn)
	if err != nil {
		fmt.Println(fmt.Errorf("execute statement: %v", err))
		return "", fmt.Errorf("internal error")
	}

	return fmt.Sprintf("Updated book %v written by %v with ISBN: %v", book.Title, book.Author, book.Isbn), nil
}

// GetBook returns one or more book from the database based on supplied filters
func (s *MySQLHandler) GetBook(filters *apis.FilterChain) (books []apis.Book, err error) {
	prepare, query := filters.SQLStatement()
	qs := fmt.Sprintf("SELECT * FROM books WHERE %s", prepare)

	fmt.Println(qs)

	for _, q := range query {
		fmt.Printf("%v\n", q)
	}

	stmt, err := s.db.Prepare(qs)

	if err != nil {
		fmt.Println(fmt.Errorf("prepare statement: %v", err))
		return nil, fmt.Errorf("internal error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(query...)
	if err != nil {
		fmt.Println(fmt.Errorf("execute statement: %v", err))
		return nil, fmt.Errorf("internal error")
	}
	defer rows.Close()

	for rows.Next() {
		book := apis.Book{}

		err = rows.Scan(&book.Title, &book.Isbn, &book.Author, &book.PublishedDate, &book.Edition, &book.Description, &book.Genre)
		if err != nil {
			fmt.Println(fmt.Errorf("scan row: %v", err))
			return nil, fmt.Errorf("internal error")
		}

		books = append(books, book)
	}

	return books, nil
}
