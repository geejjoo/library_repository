package models

type AuthorDTO struct {
	ID        int       `json:"id" db:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string    `json:"firstname" db:"firstname" db_type:"text"`
	LastName  string    `json:"lastname" db:"lastname" db_type:"text"`
	Books     []BookDTO `json:"books" db:"books" db_type:"json"`
}

func (a *AuthorDTO) TableName() string {
	return "authors"
}

func (a *AuthorDTO) GetID() int {
	return a.ID
}
func (a *AuthorDTO) GetFirstName() string {
	return a.FirstName
}
func (a *AuthorDTO) GetLastName() string {
	return a.LastName
}
func (a *AuthorDTO) GetBooks() []BookDTO {
	return a.Books
}

func (a *AuthorDTO) SetID(id int) {
	a.ID = id
}
func (a *AuthorDTO) SetFirstName(firstName string) {
	a.FirstName = firstName
}
func (a *AuthorDTO) SetLastName(lastName string) {
	a.LastName = lastName
}
func (a *AuthorDTO) SetBooks(books ...BookDTO) {
	a.Books = append(a.Books, books...)
}
