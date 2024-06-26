package postgres

import (
	"database/sql"
	"fmt"
	pb "model/genproto/generator"
)

type LibraryRepo struct {
	Db *sql.DB
}

func NewLibrayRepo(db *sql.DB) *LibraryRepo {
	return &LibraryRepo{db}
}

// kitob yaratish
func (l *LibraryRepo) ADDBOOK(book *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	fmt.Println("1.1.1.1")
	_, err := l.Db.Exec("insert into book(title, author, year_published) values($1, $2, $3)", book.Title, book.Author, book.YearPublished)
	if err != nil {
		fmt.Println(err, "2-2-2-2")
		return nil, err
	}
	resq := pb.AddBookResponse{}
	row := l.Db.QueryRow("select book_id from book")
	err = row.Scan(&resq)
	if err != nil {
		fmt.Println(err, "3-3-3-3")
		return nil, err
	}
	fmt.Println("4.4.4.4")
	return &resq, nil
}

// Id si orqali ma`lumotini olib kelish 
func (l *LibraryRepo) SearchBook(q *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
	rows, err := l.Db.Query("select book_id, title, author, year_published from book where book_id=$1", q.Query)
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


