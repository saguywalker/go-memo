package memo

import (
	"context"

	"github.com/saguywalker/go-memo/model"
)

// Usecase represent the memo's usecase
type Usecase interface {
	Fetch(context.Context) ([]*model.Note, error)
	GetByID(context.Context, []byte) (*model.Note, error)
	Store(context.Context, *model.Note) error
	Update(context.Context, *model.Note) error
}
