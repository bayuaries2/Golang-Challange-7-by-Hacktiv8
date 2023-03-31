package repository

import (
	"Challange-7/model"
	"Challange-7/repository/query"
	"fmt"
)

// handler>service>repo

// interface book
type BookRepo interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetBooks() (res []model.Book, err error)
	GetBookById(id int) (res model.Book, err error)
	UpdateBook(in model.Book, id int) (res string, err error)
	DeleteBook(id int) (res string, err error)
}

func (r Repo) CreateBook(in model.Book) (res model.Book, err error) {
	err = r.db.QueryRow(
		query.AddBook,
		in.Title,
		in.Author,
		in.Description,
	).Scan(
		&res.ID,
		&res.Title,
		&res.Author,
		&res.Description,
	)
	if err != nil {
		return res, err
	}

	// tidak error
	return res, nil
}

func (r Repo) GetBooks() (res []model.Book, err error) {
	rows, err := r.db.Query(query.GetBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var temp = model.Book{}
		err = rows.Scan(&temp.ID, &temp.Title, &temp.Author, &temp.Description)

		if err != nil {
			return res, err
		}

		res = append(res, temp)
	}

	return res, nil
}

func (r Repo) GetBookById(id int) (res model.Book, err error) {

	err = r.db.QueryRow(
		query.GetBookById,
		id,
	).Scan(
		&res.ID,
		&res.Title,
		&res.Author,
		&res.Description,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UpdateBook(in model.Book, id int) (res string, err error) {

	temp, err := r.db.Exec(
		query.UpdateBook,
		id,
		in.Title,
		in.Author,
		in.Description,
	)

	if err != nil {
		return res, err
	}

	count, err := temp.RowsAffected()
	if err != nil {
		return res, err
	}

	res = fmt.Sprint("Rows Affected ", count)

	return res, nil
}

func (r Repo) DeleteBook(id int) (res string, err error) {

	temp, err := r.db.Exec(query.DeleteBook, id)
	if err != nil {
		return res, err
	}

	count, err := temp.RowsAffected()
	if err != nil {
		return res, err
	}
	res = fmt.Sprint("Rows Affected ", count)

	return res, nil
}
