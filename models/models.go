package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func Init() {
	ConnectDatabase()
}

func ConnectDatabase() error {
	var err error
	db, err = gorm.Open("sqlite3", "/data/db.db")
	if err != nil {
		return fmt.Errorf("error in connectDatabase(): %v", err)
	}
	db.AutoMigrate(&Event{}) // ok for development, suggest using Goblin for more complex requirements.
	db.LogMode(true)
	return nil
}
