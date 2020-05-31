package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	. "blockChain/BLC"
	_ "blockChain/database"
)

type Message struct {
	Data string
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	data := r.Form.Get("data")

	chain := GetChain()
	newBlock, err := chain.Add(nil)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, data)
		return
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock.LastBlock)

}

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	chain := GetBlockList()

	bytes, err := json.MarshalIndent(chain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}

func run() error {
	mux := makeMuxRouter()
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

func main() {

	go func() {
		//chain := GetChain()
		//chain.Db.Read([]byte("l"))
		//spew.Dump(GetChain)

	}()
	run()

}
