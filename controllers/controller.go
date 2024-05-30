package controllers

import (
	"apidal/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {

	var dalam []models.Dalam

	models.DB.Find(&dalam)

	return c.Status(fiber.StatusOK).JSON(dalam)

}

func Show(c *fiber.Ctx) error {

	id := c.Params("id")

	var dalam models.Dalam

	if err := models.DB.Model(&dalam).Where("id = ?", id).First(&dalam).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Kamar dengan id " + id + " tidak ditemukan",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server sedang mengalami gangguan",
		})
	}

	return c.JSON(dalam)
}

func Create(c *fiber.Ctx) error {

	var dalam models.Dalam

	if err := c.BodyParser(&dalam); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	
	dalam.Created = time.Now()
	dalam.Updated = time.Now()

	if err := models.DB.Create(&dalam).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Antrian berhasil dibuat",
	})
}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var dalam models.Dalam

	if err := c.BodyParser(&dalam); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	dalam.Updated = time.Now()

	if models.DB.Where("id = ?", id).Updates(&dalam).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Kamar dengan id " + id + " tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diperbaharui",
	})
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")

	var dalam models.Dalam

	if models.DB.Where("id = ?", id).Delete(&dalam).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Kamar dengan id " + id + " tidak dapat dihapus",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil menghapus data kamar",
	})
}