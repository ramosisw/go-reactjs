package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ramosisw/go-reactjs/frontend"
	"github.com/ramosisw/go-reactjs/handler"
	"github.com/ramosisw/go-reactjs/models"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //sqlite3
)

const assetPrefix = "build"

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
		fmt.Printf("request URL Path %v \n", r.URL.Path)
		if strings.HasPrefix(r.URL.Path, "/_api") {
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Content-Type", "application/json")
			if r.Method == "OPTIONS" {
				w.Header().Add("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
			}
		}
		next.ServeHTTP(w, r)
	})
}
