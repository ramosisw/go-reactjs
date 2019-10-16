/*
* MIT License
*
* Copyright (c) 2019 Julio C. Ramos
*
 */

//Package models include db conection and todo model to persist with database
package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

//Init the connection with database
func Init() {
	ConnectDatabase()
}

//ConnectDatabase Connect to database, read/create sqlite file db
func ConnectDatabase() error {
	var err error
	os.MkdirAll("/data", os.ModePerm)
	db, err = gorm.Open("sqlite3", "/data/todo-go-reactjs.db")
	if err != nil {
		return fmt.Errorf("error in connectDatabase(): %v", err)
	}
	db.AutoMigrate(&Todo{})
	db.LogMode(true)
	return nil
}
