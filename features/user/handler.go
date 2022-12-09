package user

import (
	"ddrag23/gae-soal/database"
	"ddrag23/gae-soal/model"
	"ddrag23/gae-soal/utils"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	db := database.DB
	var users []model.User
	// d.DB.Find(&User{});
	db.Preload("Role").Find(&users)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "succes", "message": "ok", "data": users})
}

func Store(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}
	log.Println(user)
	roleId, err := strconv.ParseUint(c.FormValue("role_id"), 10, 32)
	if err != nil {
		log.Println(err.Error())
	}
	user.RoleId = uint(roleId)
	user.Password = hash
	db := database.DB
	log.Println(user.RoleId)

	if err := db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}
	response := ResponseUser{
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Created user", "data": response})
}

func Updated(c *fiber.Ctx) error {
	request := new(UpdateUser)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	roleId, err := strconv.ParseUint(c.FormValue("role_id"), 10, 32)
	if err != nil {
		log.Println(err.Error())
	}
	request.RoleId = uint(roleId)
	var user model.User
	db := database.DB
	db.Model(&user).Where("id = ?", c.Params("id")).Updates(model.User{
		Username: request.Username,
		Email:    request.Email,
		RoleId:   request.RoleId,
		Name:     request.Name,
	})
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Updated user", "data": request})
}

func ChangePassword(c *fiber.Ctx) error {
	request := new(ChangePasswordUser)
	id := c.Params("id")
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	var user model.User
	db := database.DB
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Data user not found", "data": nil})
	}
	log.Println(c.FormValue("old_password"))
	if !utils.CheckPasswordHash(c.FormValue("old_password"), user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid old password", "data": nil})
	}

	hash, err := utils.HashPassword(request.NewPassword)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}
	user.Password = hash
	db.Save(&user)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Change Password Success", "data": nil})

}

func Destroy(c *fiber.Ctx) error {
	id := c.Params("id")
	var user model.User

	db := database.DB
	result := db.Where("id = ?", id).Delete(&user)
	// log.Println(result.RowsAffected)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Data user not found", "data": nil})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Deleted user", "data": result.RowsAffected})
}
