package memo

import (
	"context"

	"github.com/saguywalker/go-memo/model"
)

// Repository represent the memo's repository contract
type Repository interface {
	Fetch(context.Context) ([]*model.Note, error)
	GetByID(context.Context, []byte) (*model.Note, error)
	Store(context.Context, *model.Note) error
	Update(context.Context, *model.Note) error
}
