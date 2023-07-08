package repo

import (
	"context"
	"fmt"
	"grpc-kvado/cmd/library/app"
)

func (r *Repo) GetBookByAuthor(ctx context.Context, id int) (out *[]app.Book, err error) {

	const query = `select creators.id, creators.name from books
join books_creators ON books_creators.book_id = books.id
join creators ON creators.id = books_creators.creator_id
 where books.id = $1`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var books []app.Book

	for rows.Next() {
		b := &book{}

		err = rows.Scan(&b.id, &b.name)
		if err != nil {
			fmt.Println(err)

			continue
		}

		books = append(books, *b.convert())

	}
	return &books, nil
}

func (b *book) convert() *app.Book {
	return &app.Book{
		Id:   b.id,
		Name: b.name,
	}
}
