/*
* MIT License
*
* Copyright (c) 2019 Julio C. Ramos
*
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ramosisw/todo-go-reactjs/frontend"
	handler "github.com/ramosisw/todo-go-reactjs/handlers"
	"github.com/ramosisw/todo-go-reactjs/models"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //sqlite3
)

const assetPrefix = "frontend/build/"

func assetFS() *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     frontend.Asset,
		AssetDir:  frontend.AssetDir,
		AssetInfo: frontend.AssetInfo,
		Prefix:    assetPrefix,
	}
}

func main() {
	models.Init()
	router := mux.NewRouter().StrictSlash(true)

	router.Use(commonMiddleware)
	router.HandleFunc("/_api/todo", handler.PostTodo).Methods("POST")
	router.HandleFunc("/_api/todo", handler.GetTodos).Methods("GET")
	router.HandleFunc("/_api/todo/{id}", handler.PutTodo).Methods("PUT")
	router.HandleFunc("/_api/todo/{id}", handler.GetTodo).Methods("GET")

	router.PathPrefix("/").Handler(
		http.FileServer(assetFS()),
	)

	fmt.Println("GO REST server running on http://localhost ")
	log.Fatal(http.ListenAndServe(":80", router))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s URL Path %v \n", r.Method, r.URL.Path)
		if strings.HasPrefix(r.URL.Path, "/_api") {
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Content-Type", "application/json")
			if r.Method == "OPTIONS" {
				fmt.Println("### OPTIONS REQUESTED")
				w.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE") // You can add more headers here if needed
				w.Header().Add("Access-Control-Allow-Headers", "Authorization")       // You can add more headers here if needed
				w.WriteHeader(http.StatusOK)
			} else {
				next.ServeHTTP(w, r)
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
