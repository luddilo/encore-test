package todo

import (
	"context"
	"encore.dev/storage/sqldb"
)

type CreateParams struct {
	Description string
}

type CreateResponse struct {
	ID int
}

// Create creates a new todo item.
//
// encore:api public
func Create(ctx context.Context, params *CreateParams) (*CreateResponse, error) {
	var id int
    err := sqldb.QueryRow(ctx, `
        INSERT INTO todo (description)
        VALUES ($1)
        RETURNING id
    `, params.Description).Scan(&id)
    if err != nil {
        return nil, err
    }
	return &CreateResponse{ID: id}, nil
}