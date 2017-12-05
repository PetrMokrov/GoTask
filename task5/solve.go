package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type SUrl struct {
	Url string `json:"url"`
}

type SKey struct {
	Key string `json:"key"`
}

var (
	Urls       = make(map[int]string)
	currentKey = 0
)

func NotKeyHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var newUrl SUrl
	err := json.NewDecoder(r.Body).Decode(&newUrl)
	if err != nil {
		panic(err)
	}
	Urls[currentKey] = newUrl.Url
	njAnswer := SKey{Key: strconv.Itoa(currentKey)}
	jAnswer, err := json.Marshal(njAnswer)
	if err != nil {
		panic(err)
	}
	currentKey++
	rw.WriteHeader(http.StatusOK)
	rw.Write(jAnswer)
}

func KeyHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["key"])
	if err != nil {
		panic(err)
	}
	if url, ok := Urls[key]; ok {
		http.Redirect(rw, r, url, http.StatusMovedPermanently)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", NotKeyHandler).Methods("POST")
	router.HandleFunc("/{key}", KeyHandler)
	http.ListenAndServe(":8082", router)
}
