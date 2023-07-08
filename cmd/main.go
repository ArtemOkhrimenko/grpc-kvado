package main

import (
	"context"
	"database/sql"
	app "grpc-kvado/cmd/library/app"
	repo "grpc-kvado/cmd/library/repo"
)

func main() {
	_ = context.Background()

	db, err := sql.Open("mysql", "root:password@/library-db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repo.New(db)
	_ = app.New(repository) //application

}
