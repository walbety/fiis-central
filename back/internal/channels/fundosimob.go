package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/walbety/go-fii/internal/integration/infomoney"
)

var (
	svc infomoney.Infomoney
)

type FundsController struct {
}

func GetYieldTickersFromFII(ctx *fiber.Ctx) error {
	fundo := ctx.Query("fund")
	if fundo == "" {
		ctx.Status(fiber.StatusBadRequest).SendString("must inform the funds")
		return nil
	}

	tickers, err := svc.GetYieldTickersFromFII(fundo)
	if err != nil {
		fmt.Print("error at rest/getyield", err)
		return err
	}

	err = ctx.Status(fiber.StatusOK).JSON(tickers)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		fmt.Print("error at rest/getyield", err)
		return err
	}
	return nil
}

func init() {
	svc = infomoney.New()
}
