package smsgateway

import (
	"bitbucket.org/capcom6/smsgatewaybackend/internal/config"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/infra/cli"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/infra/db"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/infra/http"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/infra/logger"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/infra/validator"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/handlers"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/models"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/repositories"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/services"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"server",
	cli.Module,
	config.Module,
	logger.Module,
	http.Module,
	validator.Module,
	handlers.Module,
	services.Module,
	repositories.Module,
	models.Module,
	db.Module,
)

func Run() {
	fx.New(
		Module,
		// fx.Invoke(cli()),
		// fx.Invoke(func(h *fiber.App) {

		// }),
		// fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		// 	return &fxevent.ZapLogger{Logger: log}
		// }),
	).Run()
}
