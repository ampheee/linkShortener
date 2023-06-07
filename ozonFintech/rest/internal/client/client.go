package client

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"ozonFintech/config"
	"ozonFintech/internal/linkService"
	"ozonFintech/internal/linkService/usecase"
)

type AppContext struct {
	App    *fiber.App
	LinkS  linkService.Link
	logger zerolog.Logger
}

func (aCtx *AppContext) CtrlGetLink(c *fiber.Ctx) error {
	aCtx.logger.Info().Msg("got " + c.Params("*") + "endpoint")
	return c.SendString(c.Params("*"))
}
func (aCtx *AppContext) CtrlPostLink(c *fiber.Ctx) error {
	aCtx.logger.Info().Msg("got " + c.Params("*") + "endpoint")
	return c.SendString(c.Params("*"))
}

func InitControllers(aCtx *AppContext) {
	aCtx.App.Get("/*", aCtx.CtrlGetLink)
	aCtx.App.Post("/*", aCtx.CtrlPostLink)
}

func NewClient(ctx context.Context, c config.Config) (*AppContext, error) {
	lService, err := usecase.NewLinkService(ctx, c)
	if err != nil {
		return nil, err
	}
	aCtx := &AppContext{
		App:   fiber.New(),
		LinkS: lService,
	}
	InitControllers(aCtx)
	return aCtx, nil
}
