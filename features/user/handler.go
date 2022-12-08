package user

import (
	"ddrag23/gae-soal/database"
	"ddrag23/gae-soal/model"
	"ddrag23/gae-soal/utils"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error  {
	// db := database.DB
	// d.DB.Find(&User{});
	return  c.SendString("oke")
}

func Store(c *fiber.Ctx) error{
	user := new(model.User)
	if err := c.BodyParser(user);err!= nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status" : "error", "message" : "Review your input","data":err})
	}
	hash,err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status" : "error", "message" : "Couldn't hash password","data":err})
	}
	user.Password = hash
	db := database.DB

	if err := db.Create(&user).Error; err != nil  {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status" : "error", "message" : "Couldn't create user","data":err})
	}
	response := ResponseUser{
		Username: user.Username,
		Name: user.Name,
		Email: user.Email,
	} 
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status":"success","message":"Created user","data" :response})
}