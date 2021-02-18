package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	rotas := mux.NewRouter().StrictSlash(true)

	rotas.HandleFunc("/", getAll).Methods("GET")

	rotas.HandleFunc("/people", create).Methods("POST")

	var port = ":3000"

	fmt.Println("Servidor rodando na porta ", port)
	log.Fatal(http.ListenAndServe(port, rotas))

}

// Person : Classe responsavel pelos seres sayajins da api
type Person struct {
	Name string
}

var people = []Person{

	Person{Name: "Kakaroto"},

	Person{Name: "Vegeta"},
}

func getAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json; charset=UTF-8")

	var p Person

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &p); err != nil {
		w.Header().Set("Content-Type", "Application/json; charset=UTF-8")
		w.WriteHeader(422)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}

	}

	json.Unmarshal(body, &p)

	people = append(people, p)

	w.Header().Set("Content-Type", "Application/json charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		panic(err)
	}

}
