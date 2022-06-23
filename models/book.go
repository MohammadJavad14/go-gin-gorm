package models

type Book struct {
	ID uint `json:"id" gorm:"primary_key"`
	BookInfo
}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	BookInfo
}

type BookInfo struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
