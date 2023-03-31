package handler

import (
	"Challange-7/helper"
	"Challange-7/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) CreateBook(c *gin.Context) {
	in := model.Book{}

	err := c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// call service
	res, err := h.app.CreateBook(in)
	if err != nil {
		helper.InternalServiceError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) GetBooks(c *gin.Context) {

	res, err := h.app.GetBooks()
	if err != nil {
		helper.InternalServiceError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) GetBookById(c *gin.Context) {
	bookId := c.Param("id")
	id, err := strconv.Atoi(bookId)
	if err != nil {
		helper.InternalServiceError(c, err.Error())
		return
	}

	res, err := h.app.GetBookById(id)
	if err != nil {
		helper.InternalServiceError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) UpdateBook(c *gin.Context) {
	bookId := c.Param("id")
	in := model.Book{}

	id, err := strconv.Atoi(bookId)
	if err != nil {
		helper.InternalServiceError(c, err.Error())
		return
	}

	eror := c.ShouldBindJSON(&in)
	if eror != nil {
		helper.BadRequest(c, eror.Error())
		return
	}

	res, err := h.app.UpdateBook(in, id)
	if err != nil {
		helper.InternalServiceError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) DeleteBook(c *gin.Context) {
	bookId := c.Param("id")
	id, err := strconv.Atoi(bookId)
	if err != nil {
		helper.InternalServiceError(c, err.Error())
		return
	}

	res, err := h.app.DeleteBook(id)
	if err != nil {
		helper.InternalServiceError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}
