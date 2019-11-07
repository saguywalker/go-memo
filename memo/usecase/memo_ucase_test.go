package usecase_test

import (
	"context"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/syndtr/goleveldb/leveldb"

	"github.com/saguywalker/go-memo/memo/repository"
	"github.com/saguywalker/go-memo/memo/usecase"
	"github.com/saguywalker/go-memo/model"
)

func TestMemoUsecase(t *testing.T) {
	dir, err := ioutil.TempDir("", "leveldb-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	db, err := leveldb.OpenFile(dir, nil)
	if err != nil {
		t.Fatal(err)
	}

	mr := repository.NewLevelDBMemoRepository(db)
	timeout := time.Duration(15 * time.Second)
	ctx := context.Background()
	mu := usecase.NewMemoUsecase(mr, timeout)

	// Fetch 1
	notes, err := mu.Fetch(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if notes != nil {
		t.Fatalf("expected nil, got %+v", notes)
	}

	// Store 1
	s1 := model.Note{
		Title:    "No.1",
		Detail:   "this is a first note",
		LastEdit: time.Now().Format(time.RFC822),
	}
	err = mu.Store(ctx, &s1)
	if err != nil {
		t.Fatal(err)
	}

	// Store 2
	s2 := model.Note{
		Title:    "No.2",
		Detail:   "this is a second note",
		LastEdit: time.Now().Format(time.RFC822),
	}
	err = mu.Store(ctx, &s2)
	if err != nil {
		t.Fatal(err)
	}

	// Get 1
	tmp, err := mu.GetByID(ctx, 2)
	if err != nil {
		t.Fatal(err)
	}
	if tmp.Id != s2.Id || tmp.Title != s2.Title || tmp.Detail != s2.Detail || !strings.EqualFold(tmp.LastEdit, s2.LastEdit) {
		t.Fatalf("expected %+v, got %+v\n", s2, tmp)
	}

	// Get2
	tmp, err = mu.GetByID(ctx, 50)
	if tmp != nil {
		t.Fatalf("expected %+v, got %+v\n", nil, tmp)
	}

	// Fetch 2
	tmpArr := make([]*model.Note, 2)
	tmpArr[0] = &s1
	tmpArr[1] = &s2
	notes, err = mu.Fetch(ctx)
	if err != nil {
		t.Fatalf("expected %+v, got %+v\n", tmpArr, notes)
	}

	// Update 1
	s1.Title = "try to edit"
	if err := mu.Update(ctx, &s1); err != nil {
		t.Fatal(err)
	}

	// Update 2
	s2.Id = 99
	if err := mu.Update(ctx, &s2); err != leveldb.ErrNotFound {
		t.Fatalf("expected %s, got %s", leveldb.ErrNotFound, err.Error())
	}
}
