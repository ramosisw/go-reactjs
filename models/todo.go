/*
* MIT License
*
* Copyright (c) 2019 Julio C. Ramos
*
 */

package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//Todo Item to add
type Todo struct {
	gorm.Model
	Title       string `gorm:"size:100" json:"title"`
	Description string `gorm:"size:500" json:"description"`
}

//Insert wrapper
func (s *Todo) Insert() error {
	if err := db.Create(s).Error; err != nil {
		return fmt.Errorf("error on Insert() %v", err)
	}

	return nil
}

//Update wrapper
func (s *Todo) Update() error {
	if err := db.Save(s).Error; err != nil {
		return fmt.Errorf("error on Update() %v", err)
	}

	return nil
}

//Read wrapper
func (s *Todo) Read(id int) error {
	if err := db.First(s, id).Error; err != nil {
		return fmt.Errorf("error in Read(%v);\n %v", id, err)
	}
	return nil
}

//ListTodos wrapper
func ListTodos() ([]Todo, error) {
	todos := []Todo{}
	if err := db.Find(&todos).Error; err != nil {
		return todos, fmt.Errorf("error in List() %v", err)
	}
	return todos, nil
}
