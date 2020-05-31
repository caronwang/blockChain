package router

import (
	. "blockChain/handler"

	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func MakeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/", HandleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", HandleWriteBlock).Methods("POST")
	muxRouter.HandleFunc("/trans", HandleTrans).Methods("POST")
	return muxRouter
}

func Run() error {
	mux := MakeMuxRouter()
	httpAddr := "8080"

	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
