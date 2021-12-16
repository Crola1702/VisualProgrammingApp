package Server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{}

	r := mux.NewRouter()
	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
