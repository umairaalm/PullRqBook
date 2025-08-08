package api

import (
	repository "UBookTsk/RepositoryPkg"
	models "UBookTsk/model"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	BookRepo repository.BookRepository
}

func NewHandler(repo repository.BookRepository) *Handler {
	return &Handler{BookRepo: repo}
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Post("/books", h.createBook)
	app.Get("/books", h.getAllBooks)
	app.Get("/books/:id", h.getBookByID)
	app.Delete("/books/:id", h.deleteBookByID)
}

// ----------------- Handlers ---------------- //

func (h *Handler) createBook(c *fiber.Ctx) error {
	var req models.AddNewBook
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input format"})
	}
	if req.Title == "" || req.Author == "" || req.Year <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Title, Author, and valid Year are required"})
	}

	book := models.Book{Title: req.Title, Author: req.Author, Year: req.Year}

	if err := h.BookRepo.Create(&book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create book"})
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

func (h *Handler) getAllBooks(c *fiber.Ctx) error {
	books, err := h.BookRepo.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve books"})
	}
	return c.JSON(books)
}

func (h *Handler) getBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	book, err := h.BookRepo.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}

func (h *Handler) deleteBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.BookRepo.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete book"})
	}
	return c.JSON(fiber.Map{"message": "Book deleted successfully"})
}
