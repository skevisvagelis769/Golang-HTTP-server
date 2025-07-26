package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var userCache = make(map[int]User)
var cacheMutex sync.RWMutex

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", 200)
}

func SetUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
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
			"Name is required",
			http.StatusBadRequest,
		)
		return
	}
	cacheMutex.Lock()
	userCache[len(userCache)+1] = user
	cacheMutex.Unlock()
	w.WriteHeader(http.StatusNoContent)

}

func GetUser(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}
	cacheMutex.RLock()
	user, ok := userCache[id]
	cacheMutex.RUnlock()
	if !ok {
		http.Error(
			w,
			"Did not find",
			http.StatusBadRequest,
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}
	cacheMutex.RLock()
	_, ok := userCache[id]
	cacheMutex.RUnlock()
	if !ok {
		http.Error(
			w,
			"Didnt find",
			http.StatusBadRequest,
		)
		return
	}
	cacheMutex.Lock()
	for i := id + 1; i < len(userCache)+1; i++ {
		userCache[i-1] = userCache[i]
	}
	delete(userCache, len(userCache))

	cacheMutex.Unlock()
	w.WriteHeader(http.StatusOK)
}
