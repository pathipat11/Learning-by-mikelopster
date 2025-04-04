package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "mydatabase"
	username     = "postgres"
	password     = "root"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&BookN{}, &Publisher{}, &Author{}, &AuthorBook{})

	publisher := Publisher{
		Details: "Pathipat Mattra",
		Name:    "pathipat",
	}
	_ = createPublisher(db, &publisher)

	// Example data for a new author
	author1 := Author{
		Name: "Pathipat",
	}
	_ = createAuthor(db, &author1)

	author2 := Author{
		Name: "Pathipat",
	}
	_ = createAuthor(db, &author2)
	// // Example data for a new book with an author
	bookN := BookN{
		Name:        "Pathipat Book",
		Author:      "FFF",
		Description: "BookN Description",
		PublisherID: publisher.ID,               // Use the ID of the publisher created above
		Authors:     []Author{author1, author2}, // Add the created author
	}
	_ = createBookWithAuthor(db, &bookN, []uint{author1.ID, author2.ID})

	// ขาเรียก

	// Example: Get a book with its publisher
	bookWithPublisher, err := getBookWithPublisher(db, 1) // assuming a book with ID 1
	if err != nil {
		// Handle error
	}

	// Example: Get a book with its authors
	bookWithAuthors, err := getBookWithAuthors(db, 1) // assuming a book with ID 1
	if err != nil {
		// Handle error
	}

	// Example: List books of a specific author
	authorBooks, err := listBooksOfAuthor(db, 1) // assuming an author with ID 1
	if err != nil {
		// Handle error
	}

	fmt.Println(bookWithPublisher)
	fmt.Println(bookWithAuthors)
	fmt.Println(authorBooks)
}

type BookN struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	PublisherID uint
	Publisher   Publisher
	Authors     []Author `gorm:"many2many:author_books;"`
}

type Publisher struct {
	gorm.Model
	Details string
	Name    string
}

type Author struct {
	gorm.Model
	Name  string
	BookN []BookN `gorm:"many2many:author_books;"`
}

type AuthorBook struct {
	AuthorID uint
	Author   Author
	BookNID  uint
	BookN    BookN
}

func createPublisher(db *gorm.DB, publisher *Publisher) error {
	result := db.Create(publisher)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func createAuthor(db *gorm.DB, author *Author) error {
	result := db.Create(author)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func createBookWithAuthor(db *gorm.DB, bookN *BookN, authorIDs []uint) error {
	// First, create the book
	if err := db.Create(bookN).Error; err != nil {
		return err
	}

	return nil
}

func getBookWithPublisher(db *gorm.DB, bookNID uint) (*BookN, error) {
	var bookN BookN
	result := db.Preload("Publisher").First(&bookN, bookNID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &bookN, nil
}

func getBookWithAuthors(db *gorm.DB, bookNID uint) (*BookN, error) {
	var bookN BookN
	result := db.Preload("Authors").First(&bookN, bookNID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &bookN, nil
}

func listBooksOfAuthor(db *gorm.DB, authorID uint) ([]BookN, error) {
	var books []BookN
	result := db.Joins("JOIN author_books on author_books.book_n_id = book_ns.id").
		Where("author_books.author_id = ?", authorID).
		Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}
