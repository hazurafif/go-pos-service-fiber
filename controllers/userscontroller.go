package controllers

import (
	"fmt"
	"net/http"

	"go-pos-service-fiber/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	var user models.Users

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var dbUser models.Users
	if err := models.DB.Where("Email = ?", user.Email).First(&dbUser).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	if user.Password != dbUser.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
	})
}

func CreateUser(c *fiber.Ctx) error {
	var user models.Users

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.Users
	if models.DB.Delete(&user, id).RowsAffected == 0 {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Cannot delete user data",
		})
	}

	return c.SendString(fmt.Sprintf("User with ID %s has been deleted", id))
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.Users // Assuming you have a struct named "User" to represent the user model
	if err := c.BodyParser(&user); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Model(&models.Users{}).Where("id = ?", id).Updates(&user).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User updated successfully",
	})
}

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.Users

	models.DB.Find(&users)

	return c.JSON(users)
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.Users
	if err := models.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Data not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Data not found",
		})

	}

	return c.JSON(user)
}
