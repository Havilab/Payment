package service

import (
	"time"

	"github.com/ekart/payments/database"
	"github.com/ekart/payments/model"
	"github.com/gofiber/fiber/v2"
)

func CreatePayment(c *fiber.Ctx) error {
	payment := new(model.Payment)

	if err := c.BodyParser(&payment); err != nil {
		c.Status(*&fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	currentTime := time.Now().Format("02-01-2006 3:4:5 PM")

	payment.CreateDate = currentTime
	result := database.PayemntDB().Create(&payment)

	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "message": "Payment is Successfully"})
}

func GetPayemnt(c *fiber.Ctx) error {
	userid := c.Params("userid")
	payment := new(model.Payment)

	database.PayemntDB().Where("userid=?", userid).Find(&payment)
	if payment.PaymentId >= 1 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "payment": payment})
	} else {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"message": "Payment not found.."})
	}
}
