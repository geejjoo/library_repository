package models

type RatingDTO struct {
	Count  int       `json:"count"`
	Author AuthorDTO `json:"author"`
}

func (r *RatingDTO) GetCount() int {
	return r.Count
}
func (r *RatingDTO) GetAuthor() AuthorDTO {
	return r.Author
}
func (r *RatingDTO) SetCount(count int) {
	r.Count = count
}
func (r *RatingDTO) SetAuthor(author AuthorDTO) {
	r.Author = author
}
