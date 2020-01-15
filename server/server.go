package server


import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	
)

// NewServer create a new server
func NewServer() *negroni.Negroni {
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx)
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router) {
	mx.HandleFunc("/api/kaggle",redirectHandler).Methods("GET")
}