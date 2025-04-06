package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "postgres"   // as defined in docker-compose.yml
	password = "root"       // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

func main() {

	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&ExampleModel{})

	c := cron.New(cron.WithSeconds())

	c.AddFunc("*/10 * * * * *", func() {
		fmt.Println("Hello World every 10 second")
		task(db)
	})

	c.Start()

	fmt.Println(db)

	select {}
}
