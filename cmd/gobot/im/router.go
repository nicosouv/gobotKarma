package im

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET", "POST").Path("/event").HandlerFunc(eventHandler)

	return router
}


func eventHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
}