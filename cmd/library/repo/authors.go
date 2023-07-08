package repo

import (
	"context"
	"fmt"
	"grpc-kvado/cmd/library/app"
)

func (r *Repo) GetAuthorByBook(ctx context.Context, id int) (out *[]app.Author, err error) {
	const query = `select books.id, books.name from creators
join books_creators ON books_creators.creator_id = creators.id
join books ON books.id = books_creators.books_id
 where creator.id = $1`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var authors []app.Author

	for rows.Next() {
		a := &author{}

		err = rows.Scan(&a.id, &a.name)
		if err != nil {
			fmt.Println(err)

			continue
		}

		authors = append(authors, *a.convert())

	}
	return &authors, nil
}

func (a *author) convert() *app.Author {
	return &app.Author{
		Id:   a.id,
		Name: a.name,
	}
}
