package rest

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	controller "github.com/walbety/go-fii/internal/channels"
)

type Response struct {
	Code    string
	Message string
}

func getResponse(ctx *fiber.Ctx) error {
	response := Response{
		"code 000",
		"message aaa",
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func create(ctx *fiber.Ctx) error {
	body := new(Response)
	err := ctx.BodyParser(body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}

	fmt.Printf("POST received: %s", body)

	return ctx.Status(fiber.StatusOK).JSON(body)

}

func Start() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(requestid.New())
	initializeEndpoints(app)
	err := app.Listen(":8080")
	if err != nil {
		fmt.Print(err)
	}
}

func initializeEndpoints(app *fiber.App) {
	app.Get("/fund", controller.GetYieldTickersFromFII)
	app.Post("/a", create)
}
