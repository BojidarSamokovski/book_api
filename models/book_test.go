package models

import (
	"testing"
)

func TestBook(t *testing.T) {
	book := Book{ID: 1, Title: "Test Book", ISBN: "0000", Author: "Test Author", Year: 2020}
	if book.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", book.ID)
	}
	if book.Title != "Test Book" {
		t.Errorf("Expected Title to be 'Test Book', got %s", book.Title)
	}
	if book.ISBN != "0000" {
		t.Errorf("Expected ISBN to be '0000', got %s", book.ISBN)
	}
	if book.Author != "Test Author" {
		t.Errorf("Expected Author to be 'Test Author', got %s", book.Author)
	}
	if book.Year != 2020 {
		t.Errorf("Expected Year to be 2020, got %d", book.Year)
	}
}
