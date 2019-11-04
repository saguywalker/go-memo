package usecase

import (
	"context"
	"time"

	"github.com/saguywalker/go-memo/memo"
	"github.com/saguywalker/go-memo/model"
)

type memoUsecase struct {
	memoRepo       memo.Repository
	contextTimeout time.Duration
}

func NewMemoUsecase(mr memo.Repository, timeout time.Duration) memo.Usecase {
	return &memoUsecase{
		memoRepo:       mr,
		contextTimeout: timeout,
	}
}

func (m *memoUsecase) Fetch(c context.Context) ([]*model.Note, error) {
	return nil, nil
}

func (m *memoUsecase) GetByID(c context.Context, id []byte) (*model.Note, error) {
	return nil, nil
}

func (m *memoUsecase) Store(c context.Context, note *model.Note) error {
	return nil
}

func (m *memoUsecase) Update(c context.Context, note *model.Note) error {
	return nil
}
