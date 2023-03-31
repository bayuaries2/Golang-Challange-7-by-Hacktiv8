package service

import (
	"Challange-7/model"
)

type BookService interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetBooks() (res []model.Book, err error)
	GetBookById(id int) (res model.Book, err error)
	UpdateBook(in model.Book, id int) (res string, err error)
	DeleteBook(id int) (res string, err error)
}

func (s *Service) CreateBook(in model.Book) (res model.Book, err error) {
	// call repository
	res, err = s.repo.CreateBook(in)
	if err != nil {
		return res, err
	}

	return res, nil
}
func (s *Service) GetBooks() (res []model.Book, err error) {
	// var books []model.Book
	res, err = s.repo.GetBooks()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) GetBookById(id int) (res model.Book, err error) {

	res, err = s.repo.GetBookById(id)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) UpdateBook(in model.Book, id int) (res string, err error) {

	res, err = s.repo.UpdateBook(in, id)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) DeleteBook(id int) (res string, err error) {

	res, err = s.repo.DeleteBook(id)
	if err != nil {
		return res, err
	}

	return res, nil
}
