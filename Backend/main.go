package main

import (
	"backend/VPANodes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var variables = VPANodes.VariableManager{SystemVariables: make(map[string]VPANodes.Variable)}

var nodes = make(map[string]interface{})

/* ***************************************************************************
 * 						Server http handling methods						 *
 * ***************************************************************************/
type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func (a *api) Router() http.Handler {
	return a.router
}

func NewServer() Server {
	a := &api{}

	r := mux.NewRouter()

	r.HandleFunc("/", a.mainPage).Methods(http.MethodGet)
	r.HandleFunc("/nodes", a.fetchNodes).Methods(http.MethodGet)
	r.HandleFunc("/nodes/create", a.createNode).Methods(http.MethodPost)

	a.router = r
	return a
}

/* ***************************************************************************
 * 						Server http handling methods						 *
 * ***************************************************************************/

type Node struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	s := NewServer()
	nodes["1"] = Node{Id: "1", Name: "Node 1"}
	nodes["2"] = Node{Id: "2", Name: "Node 2"}
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
	fmt.Println("Server started")
}

func (a *api) fetchNodes(w http.ResponseWriter, r *http.Request) {
	jsonNodes, _ := json.Marshal(nodes)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonNodes)
}

func (a *api) createNode(w http.ResponseWriter, r *http.Request) {
	var node Node
	json.NewDecoder(r.Body).Decode(&node)
	nodes[node.Id] = node
	jsonNodes, _ := json.Marshal(nodes)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonNodes)
}

func (a *api) mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(w, nil)
}
