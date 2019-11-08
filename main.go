package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/syndtr/goleveldb/leveldb"

	_memoHttpDeliver "github.com/saguywalker/go-memo/memo/delivery/http"
	_memoRepo "github.com/saguywalker/go-memo/memo/repository"
	_memoUcase "github.com/saguywalker/go-memo/memo/usecase"
)

func main() {
	dbDir := flag.String("dbdir", "memodb", "a db path for storing memo")
	flag.Parse()

	// Open leveldb from dbDir path
	db, err := leveldb.OpenFile(*dbDir, nil)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer db.Close()

	// MemoRepo for communicating with database
	memoRepo := _memoRepo.NewLevelDBMemoRepository(db)
	// MemoUcase for business logic
	memoUcase := _memoUcase.NewMemoUsecase(memoRepo, time.Duration(15*time.Second))

	// Define a router with paths
	router := mux.NewRouter()
	_memoHttpDeliver.NewMemoHandler(router.PathPrefix("/api").Subrouter(), memoUcase)

	// AllowAll in CORS ***Only in Development Step***
	cors := cors.AllowAll()
	corsHandler := cors.Handler(router)

	duration := 15 * time.Second
	srv := &http.Server{
		Handler:      corsHandler,
		Addr:         "127.0.0.1:3000",
		ReadTimeout:  duration,
		WriteTimeout: duration,
	}

	log.Println("serving api at http://127.0.0.1:3000")
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	// Graceful Shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
