package funcs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	Name string `json:"name"`
}

var UserCache = make(map[int]User)
var cacheMutex sync.RWMutex

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")

}

func GetUser(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)
		return
	}
	cacheMutex.RLock()
	user, ok := UserCache[id]
	cacheMutex.RUnlock()

	if !ok {
		http.Error(
			w,
			"Didnt find",
			http.StatusNotFound,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	j, err := json.Marshal(user)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func SetUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	log.Print("received ", user)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}
	if user.Name == "" {
		http.Error(
			w,
			"name must not be null",
			http.StatusBadRequest,
		)
		return
	}
	cacheMutex.Lock()
	UserCache[len(UserCache)+1] = user
	cacheMutex.Unlock()
	w.WriteHeader(http.StatusOK)
}
