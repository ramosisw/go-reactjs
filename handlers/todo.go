/*
* MIT License
*
* Copyright (c) 2019 Julio C. Ramos
*
 */

//Package handler include todo handlers to common REST methods
package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ramosisw/todo-go-reactjs/models"
)

//PostTodo create a todo
func PostTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: PostTodo")
	var newTodo models.Todo
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the todo title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newTodo)
	if err := newTodo.Insert(); err != nil {
		fmt.Fprintf(w, "Error on insert %v", newTodo)
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTodo)
	}
}

//GetTodo Returns a todo by ID
func GetTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetTodo")
	todoID, _ := strconv.Atoi(mux.Vars(r)["id"])

	todo := &models.Todo{}
	if err := todo.Read(todoID); err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(todo)
	}
}

//PutTodo Update a todo by ID
func PutTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: PutTodo")
	todoID, _ := strconv.Atoi(mux.Vars(r)["id"])

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the todo title and description only in order to update")
	}

	todo := &models.Todo{}
	if err := todo.Read(todoID); err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.Unmarshal(reqBody, &todo)
		todo.Update()
		json.NewEncoder(w).Encode(todo)
	}
}

//GetTodos returns all todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetTodos")
	todos, _ := models.ListTodos()
	prettyJSON, _ := json.MarshalIndent(todos, "", "    ")
	fmt.Fprint(w, string(prettyJSON))
	// json.NewEncoder(w).Encode(todos)
}
