package usecase

import (
	"context"
	"time"

	"github.com/saguywalker/go-memo/memo"
	"github.com/saguywalker/go-memo/model"
)

type memoUsecase struct {
	memoRepo memo.Repository
	contextTimeout time.Duration
}

func NewMemoUsecase(mr model.Repository, timeout time.Duration) *memoUsecase {
	return &memoUsecase{
		memoRepo: mr,
		contextTimeout: timeout,
	}

}