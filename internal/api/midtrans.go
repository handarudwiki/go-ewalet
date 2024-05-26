package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/golang-ewalet/domain"
)

type midtransApi struct {
	midtransService domain.MidtransService
	topUpService    domain.TopUpService
}

func NewMidtrans(app *fiber.App, midtransService domain.MidtransService, topUpService domain.TopUpService) {
	m := midtransApi{
		midtransService: midtransService,
		topUpService:    topUpService,
	}

	app.Post("/midtrans/payment-callback", m.paymentHandlerNotification)
}

func (m midtransApi) paymentHandlerNotification(ctx *fiber.Ctx) error {
	var notificationPayload map[string]interface{}

	err := ctx.BodyParser(&notificationPayload)

	if err != nil {
		return ctx.SendStatus(400)
	}

	orderId, exist := notificationPayload["order_id"].(string)

	if !exist {
		return ctx.SendStatus(400)
	}

	success, _ := m.midtransService.VerifyPayment(ctx.Context(), orderId)

	if success {
		_ = m.topUpService.CondfirmedTopUp(ctx.Context(), orderId)

	}

	return ctx.SendStatus(200)

}
