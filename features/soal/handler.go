package soal

import (
	"ddrag23/gae-soal/database"
	"ddrag23/gae-soal/model"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	db := database.DB
	var soal []model.Soal
	// d.DB.Find(&User{});
	db.Preload("Answers").Find(&soal)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "succes", "message": "ok", "data": soal})
}
func Store(c *fiber.Ctx) error {
	request := new(requestSoal)
	// var reqJawaban requestJawaban
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	for _, v := range c.FormValue("answers[]") {
		fmt.Println(string(v))
	}
	log.Println(c.FormValue("answers[]"))
	return c.SendStatus(200)
	// user := c.Locals("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// db := database.DB
	// err := db.Transaction(func(tx *gorm.DB) error {
	// 	soal := model.Soal{
	// 		Title:        request.Title,
	// 		TypeSoal:     request.TypeSoal,
	// 		UserId:       claims["user_id"].(string),
	// 		KunciJawaban: request.KunciJawaban,
	// 		Image:        request.Image,
	// 	}
	// 	if err := tx.Create(&soal).Error; err != nil {
	// 		return err
	// 	}
	// 	for _, val := range request.Answers {
	// 		tx.Create(&model.Jawaban{
	// 			SoalId:      soal.ID,
	// 			Content:     val.Content,
	// 			ContentType: val.ContentType,
	// 		})
	// 	}
	// 	return nil
	// })
	// if err != nil {
	// 	c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create question", "data": err})
	// }
	// return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Created question", "data": nil})
}
