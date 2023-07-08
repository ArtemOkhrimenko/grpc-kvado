package app

import "context"

type Repo interface {
	GetBookByAuthor(ctx context.Context, id int) (out *[]Book, err error)
	GetAuthorByBook(ctx context.Context, id int) (out *[]Author, err error)
}

type App struct {
	repo Repo
}

func New(repo Repo) *App {
	return &App{
		repo: repo,
	}
}
