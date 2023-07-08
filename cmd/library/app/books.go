package app

import "context"

func (a *App) GetBookByAuthor(ctx context.Context, id int) (*[]Book, error) {
	books, err := a.repo.GetBookByAuthor(ctx, id)
	if err != nil {
		return nil, err
	}

	return books, nil
}
