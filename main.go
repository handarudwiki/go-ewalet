package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/golang-ewalet/component"
	"github.com/handarudwiki/golang-ewalet/config"
	"github.com/handarudwiki/golang-ewalet/internal/api"
	"github.com/handarudwiki/golang-ewalet/internal/middleware"
	"github.com/handarudwiki/golang-ewalet/internal/repository"
	"github.com/handarudwiki/golang-ewalet/internal/service"
)

func main() {
	cnf := config.Get()
	db := component.GetDatabaseConnection()
	cacheConnection := component.GetCacheConnection()

	userRepository := repository.NewUser(db)
	transactionRepository := repository.NewTransaction(db)
	accountRepository := repository.NewAccount(db)
	notificationRepository := repository.NewNotification(db)

	emailSercice := service.NewEmail(cnf)
	userService := service.NewUser(userRepository, cacheConnection, emailSercice)
	notificationService := service.NewNotification(notificationRepository)
	transactionService := service.NewTransaction(accountRepository, cacheConnection, transactionRepository, notificationRepository)

	authMid := middleware.Authenticate(userService)

	app := fiber.New()
	api.NewAuth(app, authMid, userService)
	api.NewTransfer(app, authMid, transactionService)
	api.NewNotification(app, authMid, notificationService)
	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
