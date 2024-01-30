# bookstore_management_api
This project is a CRUD (Create, Read, Update, Delete) operation API for managing a bookstore. It is built using the Go programming language and utilizes the GORM library for interacting with a MySQL database.


![Bookstore](https://github.com/MeghanashreeC/bookstore_management_api/assets/75469288/148ae41c-8516-4034-8ae0-74bf82f0a54d)

## Functionalities Provided
This API performs the following functionalities:
1. It responds with JSON output
2. Implements the following routes
   -  GET/book/
   -  GET/book/book{id}
   -  POST/book/
   - PUT/book/book{id}
   - DELETE/book/book{id}
3. GET/book/ - Provides 200OK response for all the books available
4. GET/book/book{id} - Provides 200OK response if the requested ID is obtained else gives 404 Not Found
5. POST/book/ - Creates a neww book with 200OK response
6. PUT/book/book{id} - Updates a book with the respective ID
7. DELETE/book/book{id} - Deletes the book with the respective ID
   
## Struct Defined is
> type Book struct {
	gorm.Model         // The gorm.Model is embedded in this struct, which provides some common fields like ID, CreatedAt, UpdatedAt, and DeletedAt for managing records.
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

## How to run the application
> go build .
> go run main.go

## How to test the application
> go test ./...

## Note
Using Thunderstorm to test the endpoints




   
   
