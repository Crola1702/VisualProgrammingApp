package main

import (
	"backend/VPANodes"
	// "encoding/json"
	"fmt"
	// "log"
	"net/http"
	// "github.com/gorilla/mux"
)

var variables = VPANodes.VariableManager{SystemVariables: make(map[string]VPANodes.Variable)}
var nodes = make(map[int]*Node)
var currentNodeID int = 0

type Node struct {
	Id   int               `json:"id"`
	Type string            `json:"type"`
	Node VPANodes.Operable `json:"node"`
}

func main() {
	// s := NewServer()
	newNode("NumberNode", 2.0)
	newNode("NumberNode", 3.0)
	newNode("AritmethicNode", 1, 2, 1)
	fmt.Println(nodes[1].Node.(*VPANodes.NumberNode).In1)
	fmt.Println(nodes[2].Node.(*VPANodes.NumberNode).In1)
	fmt.Println(nodes[3].Node.(*VPANodes.AritmethicNode).Operation())
	fmt.Println("Starting server on port 3000")
	// log.Fatal(http.ListenAndServe(":3000", s.Router()))
}

/* ****************** *
 * Nodes get function *
 * ****************** */

func findNodeById(id int) Node {
	return *nodes[id]
}

/* ************************ *
 * Creating nodes functions *
 * ************************ */

func nextNodeID() int {
	currentNodeID += 1
	return currentNodeID
}

func newNode(nodeType string, params ...interface{}) {
	id := nextNodeID()
	fmt.Println("Creating node with id: ", id)
	switch nodeType {
	case "NumberNode":
		nodes[id] = &Node{
			Id:   id,
			Type: "NumberNode",
			Node: &VPANodes.NumberNode{
				Id:  id,
				In1: params[0].(float64),
			},
		}
	case "AritmethicNode":
		nodes[id] = &Node{
			Id:   id,
			Type: "AritmethicNode",
			Node: &VPANodes.AritmethicNode{
				Id:            id,
				In1:           findNodeById(params[0].(int)).Node,
				In2:           findNodeById(params[1].(int)).Node,
				OperationType: params[2].(int),
			},
		}
	}
}

/* ******** *
 * API REST *
 * ******** */

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func (a *api) Router() http.Handler {
	return a.router
}

// func NewServer() Server {
// 	a := &api{}

// 	r := mux.NewRouter()

// 	r.HandleFunc("/", a.fetchNodes).Methods(http.MethodGet)
// 	r.HandleFunc("/nodes", a.fetchNodes).Methods(http.MethodGet)
// 	r.HandleFunc("/nodes/new", a.createNode).Methods(http.MethodPost)

// 	a.router = r
// 	return a
// }
