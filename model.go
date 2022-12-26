package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "******:******@tcp(localhost:3306)/gosample1")
	if err != nil {
		panic(err)
	}
}

// bookの一覧を取得
func getBooks() (books []Book, err error) {
	rows, err := Db.Query("SELECT * FROM books")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		book := Book{}
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Category)
		if err != nil {
			return
		}
		books = append(books, book)
	}
	return
}

// 指定されたidのbookを取得
func retrieve(id int) (book Book, err error) {
	book = Book{}
	err = Db.QueryRow("SELECT * FROM books WHERE id = ?", id).Scan(&book.Id, &book.Title, &book.Author, &book.Category)
	return
}

// bookの新規作成
func (book *Book) create() (err error) {
	ins, err := Db.Prepare("INSERT INTO books (title, author, category) VALUES (?, ?, ?)")
	if err != nil {
		return
	}
	defer ins.Close()

	res, err := ins.Exec(book.Title, book.Author, book.Category)
	if err != nil {
		return
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return
	}
	book.Id = int(lastInsertID)
	return

}

//bookの更新
func (book *Book) update() (err error) {
	_, err = Db.Exec("UPDATE books SET title = ?, author = ?, category = ? WHERE id = ?",
		book.Title, book.Author, book.Category, book.Id)
	return
}

//bookの削除
func (book *Book) delete() (err error) {
	_, err = Db.Exec("DELETE FROM books WHERE id = ?", book.Id)
	return
}
