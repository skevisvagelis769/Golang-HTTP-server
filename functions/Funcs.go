package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

// comment test

var userCache = make(map[int]User) //Temporary user storage
var cacheMutex sync.RWMutex        //mutex to block other threads from accessing the userCache during read or write operations since its global

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", 200)
}

// adds a user to the userCache map
func SetUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user) //get body from request (the json we POSTed) and put it into user
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
	cacheMutex.Lock() //lock userCache and add user to it
	userCache[len(userCache)+1] = user
	cacheMutex.Unlock()
	w.WriteHeader(http.StatusNoContent)

}

// This returns the user specified by the id in the URL
func GetUser(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id")) //get the id from the part of the URL that is for the {id}
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
	j, err := json.Marshal(user) //convert the user we found earlier into json
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j) //write the json to the w writer
}

// This deletes the user based on the {id} specified in the url
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
	_, ok := userCache[id] //we access the user cache in a mutex lock.We are only reading it here so we only prevent read access
	cacheMutex.RUnlock()
	if !ok {
		http.Error(
			w,
			"Didnt find",
			http.StatusBadRequest,
		)
		return
	}
	cacheMutex.Lock() //we are reading and writing the userCache so we are completely locking access to it
	for i := id + 1; i < len(userCache)+1; i++ {
		userCache[i-1] = userCache[i]
	}
	delete(userCache, len(userCache)) //after we replace the user that we wanted to delete based on the id with the users after them so that theres continuity in the map
	//we delete the last user in the userCache (so the id equal to its length) because the last and the last - 1 users in the map will be the same!

	cacheMutex.Unlock()
	w.WriteHeader(http.StatusOK)
}
