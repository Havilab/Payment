package handler

import (
	"github.com/ekart/payments/service"
	"github.com/gofiber/fiber/v2"
)

func PaymentHandler(app *fiber.App) {

	root := app.Group("/payment")
	v := root.Group("/v1")

	v.Post("/createpayment", service.CreatePayment)
	v.Get("/getpaymentbyid/:userid", service.GetPayemnt)
}
