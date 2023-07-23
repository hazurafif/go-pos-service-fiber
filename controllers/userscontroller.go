package controllers

import (
	"fmt"

	"go-pos-service-fiber/models"

	"github.com/gofiber/fiber/v2"
)

var user models.Users

func Login(c *fiber.Ctx) error {
	return nil
}

func CreateUser(c *fiber.Ctx) error {
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	if err := models.DB.Create(&user).Error; err != nil {
		return err
	}
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := models.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return c.SendString(fmt.Sprintf("User with id %s has been deleted", id))
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	if err := models.DB.Where("id = ?", id).Updates(&user).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "User updated successfully",
	})
}

func GetAllUsers(c *fiber.Ctx) error {
	return nil
}

func GetUserByID(c *fiber.Ctx) error {
	return nil
}
