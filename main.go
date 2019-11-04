package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/gorilla/mux"

	_memoHttpDeliver "github.com/saguywalker/go-memo/memo/delivery/http"
	_memoRepo "github.com/saguywalker/go-memo/memo/repository"
	_memoUcase "github.com/saguywalker/go-memo/memo/usecase"
)

func main() {
	dbDir := flag.String("dbdir", "memodb", "a db path for storing memo")
	flag.Parse()

	db, err := badger.Open(badger.DefaultOptions(*dbDir))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer db.Close()

	memoRepo := _memoRepo.NewBadgerMemoRepository(db)
	memoUcase := _memoUcase.NewMemoUsecase(memoRepo, time.Duration(15*time.Second))

	router := mux.NewRouter()
	_memoHttpDeliver.NewMemoHandler(router.PathPrefix("/api").Subrouter(), memoUcase)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3000",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
