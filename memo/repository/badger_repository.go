package repository

import (
	"context"

	"github.com/dgraph-io/badger"
	"github.com/golang/protobuf/proto"

	"github.com/saguywalker/go-memo/memo"
	"github.com/saguywalker/go-memo/model"
)

type badgerMemoRepository struct {
	DB *badger.DB
}

func NewBadgerMemoRepository(db *badger.DB) memo.Repository {
	return &badgerMemoRepository{db}
}

func (m *badgerMemoRepository) Fetch(ctx context.Context) ([]*model.Note, error) {
	var notes []*model.Note

	err := m.DB.View(func(txn *badger.Txn) error {
		itr := txn.NewIterator(badger.DefaultIteratorOptions)
		defer itr.Close()
		for ; itr.Valid(); itr.Next() {
			valBytes, err := itr.Item().ValueCopy(nil)
			if err != nil {
				return err
			}

			var tmpNote model.Note
			if err := proto.Unmarshal(valBytes, &tmpNote); err != nil {
				return err
			}

			notes = append(notes, &tmpNote)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (m *badgerMemoRepository) GetByID(ctx context.Context, id []byte) (*model.Note, error) {
	var note model.Note

	err := m.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(id)
		if err != nil {
			return err
		}

		tmpVal, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}

		if err := proto.Unmarshal(tmpVal, &note); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	
	return &note, nil
}

func (m *badgerMemoRepository) Store(ctx context.Context, note *model.Note) error {
	noteBytes, err := proto.Marshal(note)
	if err != nil {
		return err
	}

	

	txn := m.DB.NewTransaction(true)
	defer txn.Discard()

	
	return nil
}

func (m *badgerMemoRepository) Update(ctx context.Context, note *model.Note) error {
	return nil
}
