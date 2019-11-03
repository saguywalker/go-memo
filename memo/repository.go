package memo

import (
	"context"

	"github.com/saguywalker/go-memo/model"
)

// Repository represent the memo's repository contract
type Repository interface {
	Fetch(ctx context.Context) ([]*model.Note, error)
	GetByID(ctx context.Context, id string) (*model.Note, error)
	Store(ctx context.Context, note *model.Note) error
	Update(ctx context.Context, note *model.Note) error
}
