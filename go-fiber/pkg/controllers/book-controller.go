package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ashwinpnr/golang-samples/go-fiber/pkg/models"
	"github.com/ashwinpnr/golang-samples/go-fiber/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(ctx *fiber.Ctx) error {
	log.Printf("Received GetBooks Request")
	books := models.GetAllBooks()
	log.Printf("Completed GetBooks Reques")
	ctx.JSON(books)
	return nil
}

func CreateBook(c *fiber.Ctx) error {
	log.Printf("Received Create Book Request")
	book := models.Book{}
	if err := utils.ParseBody(c, &book); err != nil {
		return err
	}

	createdBook := models.CreateBook(book)
	log.Printf("Create Book Result : %s", createdBook.ToString())
	c.JSON(createdBook)
	return nil
}

func GetBookByID(c *fiber.Ctx) error {
	bookId := c.Params("id")
	log.Printf("Received GetBook Request for Book Id : %s", bookId)
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}
	book := models.GetBookByID(ID)
	log.Printf("GetBook Result : %s", book.ToString())
	c.JSON(book)
	return nil
}

func DeleteBook(c *fiber.Ctx) error {
	bookId := c.Params("id")
	log.Printf("Received Delete Request for Book Id : %s", bookId)
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}
	book := models.DeleteBook(ID)
	log.Printf("Delete Result : %s", book.ToString())
	c.JSON(book)
	return nil
}

func UpdateBook(c *fiber.Ctx) error {
	bookId := c.Params("id")
	log.Printf("Received Delete Request for Book Id : %s", bookId)
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
	log.Printf("Delete Result : %s", updatedbook.ToString())
	c.JSON(updatedbook)
	return nil
}

func GetLive(c *fiber.Ctx) error {
	log.Printf("Liveness  Request Received")
	c.Status(http.StatusOK)
	return nil
}

func GetReady(c *fiber.Ctx) error {
	log.Printf("Readyness  Request Received")
	c.Status(http.StatusOK)
	return nil
}
