package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

// ExampleModel represents a database model
type ExampleModel struct {
	gorm.Model
	Name string
}

// task is the function to be executed on a schedule
func task(db *gorm.DB) {
	now := time.Now()
	log.Println("Task is being run...", now)

	// Example database operation
	newRecord := ExampleModel{Name: fmt.Sprintf("Record at %s", now.Format(time.RFC3339))}
	db.Create(&newRecord)
}
