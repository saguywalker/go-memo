package memo

import (
	"context"

	"github.com/saguywalker/go-memo/model"
)

// Usecase represent the memo's usecase
type Usecase interface {
	Fetch(ctx context.Context) ([]*model.Note, error)
	GetByID(ctx context.Context, id string) (*model.Note, error)
	Store(ctx context.Context, note *model.Note) error
	Update(ctx context.Context, note *model.Note) error
}
