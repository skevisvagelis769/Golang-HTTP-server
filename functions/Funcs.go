package functions

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", 200)
}

func SetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hiiii")
}
