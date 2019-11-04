package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/syndtr/goleveldb/leveldb"

	"github.com/saguywalker/go-memo/memo"
	"github.com/saguywalker/go-memo/model"
)

// MemoHandler struct define routes for memo's usecase
type MemoHandler struct {
	MHandler memo.Usecase
}

// NewMemoHandler return new MemoHandler
func NewMemoHandler(r *mux.Router, uc memo.Usecase) {
	handler := &MemoHandler{
		MHandler: uc,
	}

	r.HandleFunc("/notes", handler.Store).Methods("POST")
	r.HandleFunc("/notes", handler.Fetch).Methods("GET")
	r.HandleFunc("/notes/{id}", handler.GetByID).Methods("GET")
	r.HandleFunc("/notes/{id}", handler.Update).Methods("PUT")
}

// Store for storing note
func (m *MemoHandler) Store(w http.ResponseWriter, r *http.Request) {
	var note model.Note

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	if err := json.Unmarshal(body, &note); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := m.MHandler.Store(ctx, &note); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Memo: %d (%s) is created", note.Id, note.Title)))
}

// Fetch for retriving all notes
func (m *MemoHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	notes, err := m.MHandler.Fetch(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	notesBytes, err := json.Marshal(notes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(notesBytes)
}

// GetByID for retriving a note from id
func (m *MemoHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	vars := mux.Vars(r)
	noteID, ok := vars["id"]
	if !ok {
		http.Error(w, "missing note id", http.StatusBadRequest)
		return
	}

	note, err := m.MHandler.GetByID(ctx, []byte(noteID))
	if err != nil {
		if err == leveldb.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	noteByte, err := json.Marshal(note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(noteByte)
}

// Update for updating a note detail from id
func (m *MemoHandler) Update(w http.ResponseWriter, r *http.Request) {
	var note model.Note

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	if err := json.Unmarshal(body, &note); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := m.MHandler.Update(ctx, &note); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Memo: %d (%s) is created", note.Id, note.Title)))
}
