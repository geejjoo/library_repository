package main

import "testing"

func TestRatingList(t *testing.T) {
	Do("GET", "http://localhost:8080/rating", nil)
}

func TestRentBook(t *testing.T) {
	Do("GET", "http://localhost:8080/rent?bookID=1&userID=1", nil)
}

func TestReturnBook(t *testing.T) {
	Do("GET", "http://localhost:8080/return?bookID=1&userID=1", nil)
}
