package repository

import (
	models "UBookTsk/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *models.Book) error
	FindAll() ([]models.Book, error)
	FindByID(id string) (*models.Book, error)
	Delete(id string) error
}

type BookRepositoryImpl struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{DB: db}
}

func (r *BookRepositoryImpl) Create(book *models.Book) error {
	return r.DB.Create(book).Error
}

func (r *BookRepositoryImpl) FindAll() ([]models.Book, error) {
	var books []models.Book
	err := r.DB.Find(&books).Error
	return books, err
}

func (r *BookRepositoryImpl) FindByID(id string) (*models.Book, error) {
	var book models.Book
	err := r.DB.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepositoryImpl) Delete(id string) error {
	return r.DB.Delete(&models.Book{}, id).Error
}
