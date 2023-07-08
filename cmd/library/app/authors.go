package app

import (
	"context"
)

func (a *App) GetAuthorByBook(ctx context.Context, id int) (*[]Author, error) {
	authors, err := a.repo.GetAuthorByBook(ctx, id)
	if err != nil {
		return nil, err
	}

	return authors, nil
}
