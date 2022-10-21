package controllers

import (
	"net/http"
	"strconv"

	"github.com/ashwinpnr/golang-samples/go-fiber/pkg/models"
	"github.com/ashwinpnr/golang-samples/go-fiber/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(ctx *fiber.Ctx) error {
	books := models.GetAllBooks()
	ctx.JSON(books)
	return nil
}

func CreateBook(c *fiber.Ctx) error {
	book := models.Book{}
	if err := utils.ParseBody(c, &book); err != nil {
		return err
	}

	createdBook := models.CreateBook(book)
	c.JSON(createdBook)
	return nil
}

func GetBookByID(c *fiber.Ctx) error {
	bookId := c.Params("id")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}
	book := models.GetBookByID(ID)
	c.JSON(book)
	return nil
}

func DeleteBook(c *fiber.Ctx) error {
	bookId := c.Params("id")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}
	book := models.DeleteBook(ID)
	c.JSON(book)
	return nil
}

func UpdateBook(c *fiber.Ctx) error {
	bookId := c.Params("id")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}
	booktoupdate := models.Book{}
	if err := utils.ParseBody(c, &booktoupdate); err != nil {
		return err
	}
	updatedbook := models.UpdateBook(ID, booktoupdate)
	c.JSON(updatedbook)
	return nil
}
