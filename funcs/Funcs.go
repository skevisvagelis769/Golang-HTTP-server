package funcs

import (
	"encoding/json"
	//"fmt"

	"net/http"
	"strconv"
	"sync"
)

type User struct {
	Name string `json:"name"`
}

var UserCache = make(map[int]User)
var cacheMutex sync.RWMutex

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
			"name must not be nil",
			http.StatusBadRequest,
		)
		return
	}
	cacheMutex.Lock()
	UserCache[len(UserCache)+1] = user
	cacheMutex.Unlock()
	w.WriteHeader(http.StatusOK)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idnum, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		var user User
		var found bool = false
		var foundID int = 0
		for i, users := range UserCache {
			if users.Name == id {
				found = true
				foundID = i
				break
			}

		}
		if found {
			w.Header().Set("Content-type", "application/json")
			user.Name = "Your specified name " + id + " was found and they have an id: " + strconv.Itoa(foundID)
			js, err := json.Marshal(user)
			if err != nil {
				http.Error(
					w,
					err.Error(),
					http.StatusBadRequest,
				)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(js)
		}
	} else {
		cacheMutex.RLock()
		user, ok := UserCache[idnum]
		cacheMutex.RUnlock()
		if !ok {
			http.Error(
				w,
				"Could not find user",
				http.StatusNotFound,
			)
			return
		}
		w.Header().Set("Content-type", "application/json")
		j, err := json.Marshal(user)
		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}

}
func SrchId(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/media/carabi/New Volume/GO/del/static/search_id.html")
}
func SrchName(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/media/carabi/New Volume/GO/del/static/search_name.html")
}
