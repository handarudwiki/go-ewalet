package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/golang-ewalet/domain"
	"github.com/handarudwiki/golang-ewalet/dto"
	"github.com/handarudwiki/golang-ewalet/internal/util"
)

type transferApi struct {
	transactionService domain.TransactionService
	factorService      domain.FactoryService
}

func NewTransfer(app *fiber.App, autMid fiber.Handler, transactionService domain.TransactionService, factorService domain.FactoryService) {

	h := transferApi{
		transactionService: transactionService,
		factorService:      factorService,
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

	user := ctx.Locals("x-user").(dto.UserData)

	err = t.factorService.ValidatePIN(ctx.Context(), dto.ValidatePinReq{
		UserID: user.ID,
		PIN:    req.PIN,
	})

	if err != nil {
		return ctx.Status(util.GetHttpStatus(err)).JSON(dto.Response{
			Message: err.Error(),
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
