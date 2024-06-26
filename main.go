package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/golang-ewalet/component"
	"github.com/handarudwiki/golang-ewalet/config"
	"github.com/handarudwiki/golang-ewalet/dto"
	"github.com/handarudwiki/golang-ewalet/internal/api"
	"github.com/handarudwiki/golang-ewalet/internal/middleware"
	"github.com/handarudwiki/golang-ewalet/internal/repository"
	"github.com/handarudwiki/golang-ewalet/internal/service"
	"github.com/handarudwiki/golang-ewalet/internal/sse"
)

func main() {
	cnf := config.Get()
	db := component.GetDatabaseConnection()
	cacheConnection := component.GetCacheConnection()

	hub := &dto.Hub{
		NotificationChannel: map[int64]chan dto.NotificationData{},
	}

	userRepository := repository.NewUser(db)
	transactionRepository := repository.NewTransaction(db)
	accountRepository := repository.NewAccount(db)
	notificationRepository := repository.NewNotification(db)
	templateRepository := repository.NewTemplate(db)
	toupRepository := repository.NewTopUp(db)
	factorRepository := repository.NewFactor(db)

	emailSercice := service.NewEmail(cnf)
	userService := service.NewUser(userRepository, cacheConnection, emailSercice)
	notificationService := service.NewNotification(notificationRepository, templateRepository, hub)
	transactionService := service.NewTransaction(accountRepository, cacheConnection, transactionRepository, notificationService)
	midtranService := service.NewMidtrasn(cnf)
	toupService := service.NewTopup(notificationService, toupRepository, midtranService, accountRepository, transactionRepository)
	serviceFactor := service.NewFactor(factorRepository)

	authMid := middleware.Authenticate(userService)

	app := fiber.New()
	api.NewAuth(app, authMid, userService)
	api.NewTransfer(app, authMid, transactionService, serviceFactor)
	api.NewNotification(app, authMid, notificationService)
	api.NewTopUpApi(app, authMid, toupService)
	api.NewMidtrans(app, midtranService, toupService)

	sse.NewNotification(app, authMid, hub)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
