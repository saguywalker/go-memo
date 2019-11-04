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

// NewMemoUsecase return a memoUsecase with memo's repository and duration
func NewMemoUsecase(mr memo.Repository, timeout time.Duration) memo.Usecase {
	return &memoUsecase{
		memoRepo:       mr,
		contextTimeout: timeout,
	}
}

func (m *memoUsecase) Fetch(c context.Context) ([]*model.Note, error) {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()
	return m.memoRepo.Fetch(ctx)
}

func (m *memoUsecase) GetByID(c context.Context, id []byte) (*model.Note, error) {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()
	return m.memoRepo.GetByID(ctx, id)
}

func (m *memoUsecase) Store(c context.Context, note *model.Note) error {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()
	return m.memoRepo.Store(ctx, note)
}

func (m *memoUsecase) Update(c context.Context, note *model.Note) error {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()
	return m.memoRepo.Update(ctx, note)
}
