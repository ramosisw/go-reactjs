package models

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Event struct {
	// gorm.Model
	Id          int    `gorm:"AUTO_INCREMENT" json:"id"`
	Title       string `gorm:"size:100" json:"title"`
	Description string `gorm:"size:500" json:"description"`
}

//Insert wrapper
func (s *Event) Insert() error {
	if err := db.Create(s).Error; err != nil {
		return fmt.Errorf("error in Insert() %v", err)
	}

	return nil
}

//Read wrapper
func (s *Event) Read(id int) error {
	if err := db.First(s, id).Error; err != nil {
		return fmt.Errorf("error in GetByID() for id %v: %v", id, err)
	}
	return nil
}

//ListEvents wrapper
func ListEvents() ([]Event, error) {
	events := []Event{}
	if err := db.Find(&events).Error; err != nil {
		return events, fmt.Errorf("error in List() %v", err)
	}
	return events, nil
}
