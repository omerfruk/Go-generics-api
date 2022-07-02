package services

import (
	"github.com/omerfruk/Go-generics-api/model"
	"gorm.io/gorm"
)

type BookService struct {
	BaseService[model.Book]
}

func NewBookService(db *gorm.DB) BookService {
	return BookService{
		BaseService[model.Book]{DB: db},
	}
}
func (b BookService) GetBookByAuthor(author string) ([]model.Book, error) {
	var book []model.Book
	err := b.DB.Model(book).Where("author = ?", author).Find(&book).Error
	return book, err
}
