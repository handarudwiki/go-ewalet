package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/golang-ewalet/domain"
	"github.com/handarudwiki/golang-ewalet/dto"
)

type transferApi struct {
	transactionService domain.TransactionService
}

func NewTransfer(app *fiber.App, autMid fiber.Handler, transactionService domain.TransactionService) {

	h := transferApi{
		transactionService: transactionService,
	}

	app.Post("transfer/inquiry", autMid, h.TransferInquiry)
	app.Post("transfer/execute", autMid, h.TransferExecute)
}

func (t transferApi) TransferInquiry(ctx *fiber.Ctx) error {
	var req dto.TransferInquiryReq

	err := ctx.BodyParser(&req)

	if err != nil {
		return ctx.Status(400).JSON(dto.Response{
			Message: "invalid body",
		})
	}

	inquiry, err := t.transactionService.TransferInquiry(ctx.Context(), req)
	if err != nil {
		return ctx.Status(400).JSON(dto.Response{
			Message: err.Error(),
		})
	}

	return ctx.Status(200).JSON(inquiry)
}

func (t transferApi) TransferExecute(ctx *fiber.Ctx) error {
	var req dto.TransferExecuteReq

	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(400).JSON(dto.Response{
			Message: "invalid body",
		})
	}

	err = t.transactionService.TransferExecute(ctx.Context(), req)
	if err != nil {
		return ctx.Status(400).JSON(dto.Response{
			Message: err.Error(),
		})
	}
	return ctx.SendStatus(200)
}
