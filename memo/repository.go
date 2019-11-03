package memo

import (
	"context"

	"github.com/saguywalker/go-memo/model"
)

// Repository represent the memo's repository contract
type Repository interface {
	Fetch(ctx context.Context) ([]*model.Note, error)
	GetByID(ctx context.Context, id string) (*model.Note, error)
	Store(context.Context, *model.Note) error
	Update(context.Context, note *model.Note) error
}