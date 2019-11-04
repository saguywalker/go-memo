package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/syndtr/goleveldb/leveldb"

	_memoHttpDeliver "github.com/saguywalker/go-memo/memo/delivery/http"
	_memoRepo "github.com/saguywalker/go-memo/memo/repository"
	_memoUcase "github.com/saguywalker/go-memo/memo/usecase"
)

func main() {
	dbDir := flag.String("dbdir", "memodb", "a db path for storing memo")
	flag.Parse()

	db, err := leveldb.OpenFile(*dbDir, nil)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer db.Close()

	memoRepo := _memoRepo.NewLevelDBMemoRepository(db)
	memoUcase := _memoUcase.NewMemoUsecase(memoRepo, time.Duration(15*time.Second))

	router := mux.NewRouter()
	_memoHttpDeliver.NewMemoHandler(router.PathPrefix("/api").Subrouter(), memoUcase)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3000",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Println("serving api at http://127.0.0.1:3000")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}
