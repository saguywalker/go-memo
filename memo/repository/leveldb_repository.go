package repository

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/syndtr/goleveldb/leveldb"

	"github.com/saguywalker/go-memo/memo"
	"github.com/saguywalker/go-memo/model"
)

type leveldbMemoRepository struct {
	DB *leveldb.DB
}

// NewLevelDBMemoRepository wrap a badger database
func NewLevelDBMemoRepository(db *leveldb.DB) memo.Repository {
	return &leveldbMemoRepository{db}
}

// LastNoteID return a last note id (1 as default)
func (m *leveldbMemoRepository) LastNoteID() uint64 {
	iter := m.DB.NewIterator(nil, nil)
	defer iter.Release()
	if ok := iter.Last(); ok {
		keyBytes := iter.Key()
		key, _ := binary.Uvarint(keyBytes)
		return key
	}

	return 1
}

// Fetch retrive all notes from database
func (m *leveldbMemoRepository) Fetch(ctx context.Context) ([]*model.Note, error) {
	var notes []*model.Note

	iter := m.DB.NewIterator(nil, nil)
	for iter.Next() {
		noteBytes := iter.Value()

		var note model.Note
		if err := proto.Unmarshal(noteBytes, &note); err != nil {
			// notes = append(notes, &note)
			return nil, err
		}

		notes = append(notes, &note)
	}

	return notes, nil
}

// GetByID retrive a note from note id
func (m *leveldbMemoRepository) GetByID(ctx context.Context, id []byte) (*model.Note, error) {
	var note model.Note

	noteBytes, err := m.DB.Get(id, nil)
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(noteBytes, &note); err != nil {
		return nil, err
	}

	return &note, nil
}

// Store store a new note
func (m *leveldbMemoRepository) Store(ctx context.Context, note *model.Note) error {
	keyID := make([]byte, 8)
	binary.PutUvarint(keyID, note.Id)
	_, err := m.DB.Get(keyID, nil)
	if err == nil {
		return errors.New("duplicate note's id")
	}

	if err != nil && err != leveldb.ErrNotFound {
		return err
	}

	noteBytes, err := proto.Marshal(note)
	if err != nil {
		return err
	}

	if err := m.DB.Put(keyID, noteBytes, nil); err != nil {
		return err
	}

	return nil
}

// Update update a note from note id
func (m *leveldbMemoRepository) Update(ctx context.Context, note *model.Note) error {
	keyID := make([]byte, 8)
	binary.PutUvarint(keyID, note.Id)
	oldKey, err := m.DB.Get(keyID, nil)
	if err != nil {
		return err
	}

	if len(oldKey) == 0 {
		return fmt.Errorf("does not exist")
	}

	noteBytes, err := proto.Marshal(note)
	if err != nil {
		return err
	}

	if err := m.DB.Put(keyID, noteBytes, nil); err != nil {
		return err
	}

	return nil
}
