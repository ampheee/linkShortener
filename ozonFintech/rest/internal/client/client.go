package client

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"ozonFintech/config"
	"ozonFintech/internal/linkService"
	"ozonFintech/internal/linkService/usecase"
	"ozonFintech/pkg/logger"
)

type AppContext struct {
	App    *fiber.App
	LinkS  linkService.Link
	logger zerolog.Logger
}

func (aCtx *AppContext) CtrlGetLink(c *fiber.Ctx) error {
	aCtx.logger.Info().Msg("got " + c.Params("*") + " endpoint")
	str, err := aCtx.LinkS.GetOriginalByAbbreviated(c.UserContext(), c.Params("*"))
	if err != nil {
		if err.Error() == "redis: nil" || err.Error() == "no rows in result set" {
			return c.Status(404).SendString("Link not found")
		} else {
			return c.Status(500).SendString("InternalServiceError")
		}
	}
	aCtx.logger.Info().Msg("endpoint processed.")
	return c.Status(200).SendString(str)
}
func (aCtx *AppContext) CtrlPostLink(c *fiber.Ctx) error {
	aCtx.logger.Info().Msg("got " + "rus.tam" + c.Params("*") + " endpoint")
	link, err := aCtx.LinkS.SaveOriginalLink(c.UserContext(), c.Params("*"))
	if err != nil {
		return c.Status(500).SendString("InternalServiceError")
	}
	aCtx.logger.Info().Msg("endpoint processed.")
	return c.SendString("rus.tam/" + link)
}

func InitControllers(aCtx *AppContext) {
	aCtx.App.Get("rus.tam/*", aCtx.CtrlGetLink)
	aCtx.App.Post("/*", aCtx.CtrlPostLink)
}

func NewClient(ctx context.Context, c config.Config) (*AppContext, error) {
	lService, err := usecase.NewLinkService(ctx, c)
	if err != nil {
		return nil, err
	}
	aCtx := &AppContext{
		App:    fiber.New(),
		LinkS:  lService,
		logger: logger.GetLogger(),
	}
	InitControllers(aCtx)
	return aCtx, nil
}
