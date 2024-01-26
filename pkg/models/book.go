package models

import (
	"github.com/MeghanashreeC/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/MeghanashreeC/go-bookstore/pkg/routes"
)

var db *gorm.DB // this is used to hold connection to the db

type Book struct {
	gorm.Model         // The gorm.Model is embedded in this struct, which provides some common fields like ID, CreatedAt, UpdatedAt, and DeletedAt for managing records.
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()        // Calls config.Connect(), which presumably establishes a connection to the database.
	db = config.GetDB()     // Retrieves the database connection
	db.AutoMigrate(&Book{}) //Calls db.AutoMigrate(&Book{}) to automatically create or update the table in the database corresponding to the Book model.
}

// Overall, this code is setting up a GORM model for a Book, establishing a connection to the database,
//and ensuring that the necessary table is created or updated in
//the database when the application starts.

// Creating all the functions I need to talk to the database
// 1. Create
func (b *Book) CreateBook() *Book { //Book is what we return of type Book
	db.NewRecord(b) // This function is used to check if the record associated with the book (b) is new (i.e., hasn't been saved to the database).
	db.Create(&b)   // This line creates a new record in the database. The &b passes a pointer to the Book struct to the Create function.
	return b
}

// 2. Get

func GetAllBooks() []Book { // I am using slice because I want to get a slice of books
	var Books []Book //This is the container that will hold the books retrieved from the database.
	db.Find(&Books)  // Uses GORM's Find method to retrieve all records from the Book table
	return Books
}

// 3. getting book by id

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

// 4. Delete Book

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}

// Update will happen in controller
