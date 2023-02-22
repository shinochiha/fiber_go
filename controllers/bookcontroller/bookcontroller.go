package bookcontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shinochiha/fiber_go/models"
	"gorm.io/gorm"
)

func GetList(c *fiber.Ctx) error {
	var books []models.Book
	models.DB.Find(&books)
	return c.Status(fiber.StatusOK).JSON(books)
}

func GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if err := models.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}
	return c.JSON(book)
}

func Create(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Create(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if models.DB.Where("id = ?", id).Updates(&book).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "book with id " + id + " not found",
		})
	}

	return c.JSON(book)
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	if models.DB.Delete(&book, id).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "book with id " + id + " not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "book with id " + id + " successfully deleted",
	})
}
