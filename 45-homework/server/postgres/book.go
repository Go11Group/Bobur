package postgres

import (
	"database/sql"
	pb "model/genproto/generator"
	"strconv"
)

type LibraryRepo struct {
	Db *sql.DB
}

func NewLibrayRepo(db *sql.DB) *LibraryRepo {
	return &LibraryRepo{db}
}

// kitob yaratish
func (l *LibraryRepo) ADDBOOK(book *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	_, err := l.Db.Exec("insert into book(title, author, year_published) values($1, $2, $3)", book.Title, book.Author, int64(book.YearPublished))
	if err != nil {
		return nil, err
	}
	resq := pb.AddBookResponse{}
	row := l.Db.QueryRow("select book_id from book")
	err = row.Scan(&resq.BookId)
	if err != nil {
		return nil, err
	}
	return &resq, nil
}

// Id si orqali ma`lumotini olib kelish
func (l *LibraryRepo) SearchBook(q *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
	id, err := strconv.Atoi(q.Query)
	if err != nil {
		return nil, err
	}

	rows, err := l.Db.Query("select book_id, title, author, year_published from book where book_id=$1", int64(id))
	if err != nil {
		return nil, err
	}
	books := pb.SearchBookResponse{}
	i := 0
	for rows.Next() {
		book := pb.Book{}
		err = rows.Scan(&book.BookId, &book.Title, &book.Author, &book.YearPublished)
		if err != nil {
			return nil, err
		}
		i++
		books.Books = append(books.Books, &book)

	}

	return &books, nil
}

// func (l *LibraryRepo) BOORROBook(q *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
// 	_, err := l.Db.Exec("insert into borrow")
// }
