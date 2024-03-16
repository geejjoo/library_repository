package models

type BookDTO struct {
	ID       int    `json:"id" db:"id" db_type:"SERIAL PRIMARY KEY"`
	Title    string `json:"title"  db:"title" db_type:"text"`
	AuthorID int    `json:"authorid" db:"authorid" db_type:"integer"`
	ReaderID int    `json:"readerid" db:"readerid" db_type:"integer"`
}

func (b *BookDTO) TableName() string {
	return "books"
}

func (b *BookDTO) GetID() int {
	return b.ID
}
func (b *BookDTO) GetTitle() string {
	return b.Title
}
func (b *BookDTO) GetAuthorID() int {
	return b.AuthorID
}

func (b *BookDTO) GetReaderID() int {
	return b.ReaderID
}

func (b *BookDTO) SetID(id int) {
	b.ID = id
}
func (b *BookDTO) SetTitle(title string) {
	b.Title = title
}
func (b *BookDTO) SetAuthorID(authorID int) {
	b.AuthorID = authorID
}

func (b *BookDTO) SetReaderID(readerID int) {
	b.ReaderID = readerID
}
