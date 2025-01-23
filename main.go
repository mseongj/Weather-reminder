package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct{
	FristName string
	LastName string
	Email string
	CreateAt time.Time
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	user := new(User)
	json.NewDecoder(r.Body).Decode(user)
}

func barHandler(w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}


func main(){
	mux := http.NewServeMux();

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/bar", barHandler)
	mux.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", mux)
}
