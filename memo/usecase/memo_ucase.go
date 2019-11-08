package usecase

import (
	"context"
	"encoding/binary"
	"time"

	"github.com/saguywalker/go-memo/memo"
	"github.com/saguywalker/go-memo/model"
)

type memoUsecase struct {
	memoRepo       memo.Repository
	contextTimeout time.Duration
	CounterID      uint64
}

// NewMemoUsecase return a memoUsecase with memo's repository and duration
func NewMemoUsecase(mr memo.Repository, timeout time.Duration) memo.Usecase {
	return &memoUsecase{
		memoRepo:       mr,
		contextTimeout: timeout,
		CounterID:      mr.LastNoteID(),
	}
}

func (m *memoUsecase) Fetch(c context.Context) ([]*model.Note, error) {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()

	return m.memoRepo.Fetch(ctx)
}

func (m *memoUsecase) GetByID(c context.Context, id uint64) (*model.Note, error) {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()

	keyID := make([]byte, 8)
	binary.PutUvarint(keyID, id)

	return m.memoRepo.GetByID(ctx, keyID)
}

func (m *memoUsecase) Store(c context.Context, note *model.Note) error {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()

	note.LastEdit = time.Now().Format(time.RFC822)
	note.Id = m.CounterID

	if err := m.memoRepo.Store(ctx, note); err != nil {
		return err
	}

	m.CounterID++
	return nil
}

func (m *memoUsecase) Update(c context.Context, note *model.Note) error {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()

	note.LastEdit = time.Now().Format(time.RFC822)

	return m.memoRepo.Update(ctx, note)
}
