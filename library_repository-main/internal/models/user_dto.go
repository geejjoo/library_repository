package models

type UserDTO struct {
	ID          int       `json:"id" db:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName   string    `json:"firstname" db:"firstname" db_type:"text"`
	LastName    string    `json:"lastname" db:"lastname" db_type:"text"`
	RentedBooks []BookDTO `json:"rentedbooks" db:"rentedbooks" db_type:"text"`
}

func (u *UserDTO) TableName() string {
	return "users"
}

func (u *UserDTO) GetID() int {
	return u.ID
}
func (u *UserDTO) GetFirstName() string {
	return u.FirstName
}
func (u *UserDTO) GetLastName() string {
	return u.LastName
}
func (u *UserDTO) GetRentedBooks() []BookDTO {
	return u.RentedBooks
}
func (u *UserDTO) SetID(id int) {
	u.ID = id
}
func (u *UserDTO) SetFirstName(firstName string) {
	u.FirstName = firstName
}
func (u *UserDTO) SetLastName(lastName string) {
	u.LastName = lastName
}
func (u *UserDTO) AddRentedBooks(books ...BookDTO) {
	u.RentedBooks = append(u.RentedBooks, books...)
}

func (u *UserDTO) DeleteRentedBooks(bookID int) {
	idx := -1
	for i, v := range u.RentedBooks {
		if v.ID == bookID {
			idx = i
			break
		}
	}
	if idx != -1 {
		u.RentedBooks = append(u.RentedBooks[:idx], u.RentedBooks[idx+1:]...)
	}
}
