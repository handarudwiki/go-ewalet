package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/golang-ewalet/domain"
	"github.com/handarudwiki/golang-ewalet/dto"
	"github.com/handarudwiki/golang-ewalet/internal/util"
)

type topUpApi struct {
	topUpService domain.TopUpService
}

func NewTopUpApi(app *fiber.App, authMid fiber.Handler, topUpService domain.TopUpService) {
	t := &topUpApi{
		topUpService: topUpService,
	}

	app.Post("/topup-initialize", authMid, t.InitializeTopUp)
}

func (t topUpApi) InitializeTopUp(ctx *fiber.Ctx) error {
	var req dto.ToUpReg

	err := ctx.BodyParser(&req)

	if err != nil {
		return ctx.SendStatus(400)
	}

	user := ctx.Locals("x-user").(dto.UserData)
	req.UserID = user.ID

	res, err := t.topUpService.InitialieTopUp(ctx.Context(), req)

	if err != nil {
		return ctx.Status(util.GetHttpStatus(err)).JSON(dto.Response{
			Message: err.Error(),
		})
	}

	return ctx.Status(200).JSON(res)
}
