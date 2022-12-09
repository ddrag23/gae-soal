package auth

import (
	"ddrag23/gae-soal/config"
	"ddrag23/gae-soal/database"
	"ddrag23/gae-soal/features/user"
	"ddrag23/gae-soal/model"
	"ddrag23/gae-soal/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)


func getByUsername(username string) (*model.User,error) {
	db := database.DB
	var user model.User
	if err := db.Where("username = ?",username).First(&user).Error; err != nil {
		// if errors.Is(err, gorm.ErrRecordNotFound) {
		// 	return nil, nil
		// }
		return nil, err
	}
	return &user, nil
}
func Login(c *fiber.Ctx) error {
	request := new(RequestLogin)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}
	var ud user.ResponseUser 
	queryUser,err := getByUsername(request.Username)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on username", "data": err.Error()})
	}
	if queryUser != nil {
		ud = user.ResponseUser{
			ID: queryUser.ID,
			Username: queryUser.Username,
			Email: queryUser.Email,
			Name: queryUser.Name,
			Role: queryUser.Role,
		}
	}
	if !utils.CheckPasswordHash(request.Password,queryUser.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = ud.Username
	claims["user_id"] = ud.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}