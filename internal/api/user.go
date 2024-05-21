package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/golang-ewalet/domain"
	"github.com/handarudwiki/golang-ewalet/dto"
	"github.com/handarudwiki/golang-ewalet/internal/util"
)

type authApi struct {
	userService domain.UserService
}

func NewAuth(app *fiber.App, authMid fiber.Handler, userService domain.UserService) {
	h := authApi{
		userService: userService,
	}

	app.Post("login", h.GenerateToken)
	app.Get("token/validate", authMid, h.ValidateToken)
	app.Post("register", h.RegisterUser)
	app.Post("validate-otp", h.ValidateOTP)
}

func (a authApi) GenerateToken(ctx *fiber.Ctx) error {
	var req dto.AuthReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(400)
	}

	token, err := a.userService.Authenticate(ctx.Context(), req)
	if err != nil {
		return ctx.SendStatus(util.GetHttpStatus(err))
	}
	return ctx.Status(200).JSON(token)
}

func (a authApi) ValidateToken(ctx *fiber.Ctx) error {
	user := ctx.Locals("x-user")
	return ctx.Status(200).JSON(user)
}
func (a authApi) RegisterUser(ctx *fiber.Ctx) error {
	var req dto.UserRegisterReq

	err := ctx.BodyParser(&req)

	if err != nil {
		return ctx.SendStatus(400)
	}

	res, err := a.userService.Register(ctx.Context(), req)

	if err != nil {
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON(res)
}

func (a authApi) ValidateOTP(ctx *fiber.Ctx) error {
	var req dto.ValidateOtpReq

	err := ctx.BodyParser(&req)

	if err != nil {
		return ctx.SendStatus(40)
	}

	err = a.userService.ValidateOTP(ctx.Context(), req)

	if err != nil {
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.SendStatus(200)
}
