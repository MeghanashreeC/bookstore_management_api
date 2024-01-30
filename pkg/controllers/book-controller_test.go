package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MeghanashreeC/go-bookstore/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestGetBook(t *testing.T) {
	// Setup
	res, err := http.NewRequest("GET", "/book", nil) // creates HTTP GET request with the path /book and no request body. //http.NewRequest returns an http.Request object and an err if any
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder() // It's used to capture the HTTP response that would be sent to the client.

	// handler
	handler := http.HandlerFunc(GetBook) //This line creates an HTTP handler using the GetBook function.
	handler.ServeHTTP(rr, res)           // This line invokes the HTTP handler with the created request (res) and records the response in the ResponseRecorder (rr).

	// Assertions
	assert.Equal(t, http.StatusOK, rr.Code)      //This ensures that the request was handled successfully.
	var book []models.Book                       //Declares a variable books to hold the decoded JSON response.
	err = json.Unmarshal(rr.Body.Bytes(), &book) //Attempts to decode the response body into a slice of models.Book using json.Unmarshal
	assert.Nil(t, err)                           // Asserts that the decoding operation didn't result in an error
}

func TestGetBookById(t *testing.T) {
	// setup
	res, err := http.NewRequest("GET", "/book/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// handler

	handler := http.HandlerFunc(GetBookById)
	handler.ServeHTTP(rr, res)

	//Assertions

	assert.Equal(t, http.StatusOK, rr.Code)
	var book models.Book
	err = json.Unmarshal(rr.Body.Bytes(), &book)
	assert.Nil(t, err)
}

func TestCreateBook(t *testing.T) {
	// setup
	bookJSON := `{"Name": "Test Book", "Author": "Test Author", "Publication": "Test Publication"}`
	res, err := http.NewRequest("POST", "/book", strings.NewReader(bookJSON))
	if err != nil {
		t.Fatal(err)
	}
	res.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// handler
	handler := http.HandlerFunc(CreateBook)
	handler.ServeHTTP(rr, res)

	// Assertions
	assert.Equal(t, http.StatusOK, rr.Code)
	var createdBook models.Book
	err = json.Unmarshal(rr.Body.Bytes(), &createdBook)
	assert.Nil(t, err)
}

func TestDeleteBook(t *testing.T) {
	// Setup
	res, err := http.NewRequest("DELETE", "/book/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Handler
	handler := http.HandlerFunc(DeleteBook)
	handler.ServeHTTP(rr, res)

	// Assertions
	assert.Equal(t, http.StatusOK, rr.Code)
	var deletedBook models.Book
	err = json.Unmarshal(rr.Body.Bytes(), &deletedBook)
	assert.Nil(t, err)
	// Additional assertions based on your application logic
}

func TestUpdateBook(t *testing.T) {
	// Setup
	updateData := `{"Name": "Updated Book", "Author": "Updated Author", "Publication": "Updated Publication"}`
	res, err := http.NewRequest("PUT", "/book/1", strings.NewReader(updateData))
	if err != nil {
		t.Fatal(err)
	}
	res.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Handler
	handler := http.HandlerFunc(UpdateBook)
	handler.ServeHTTP(rr, res)

	// Assertions
	assert.Equal(t, http.StatusOK, rr.Code)
	var updatedBook models.Book
	err = json.Unmarshal(rr.Body.Bytes(), &updatedBook)
	assert.Nil(t, err)
	// Additional assertions based on your application logic
}
