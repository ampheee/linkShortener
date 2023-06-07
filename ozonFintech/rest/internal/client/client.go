package client

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"ozonFintech/config"
	"ozonFintech/internal/linkService"
	"ozonFintech/internal/linkService/usecase"
)

type AppContext struct {
	App   *fiber.App
	LinkS linkService.Link
}

func (aCtx *AppContext) CtrlGetLink(c *fiber.Ctx) error {
	return c.SendString(c.Params("*"))
}
func (aCtx *AppContext) CtrlPostLink(c *fiber.Ctx) error {
	return c.SendString(c.Params("*"))
}

func InitControllers(aCtx *AppContext) {
	aCtx.App.Get("/*", aCtx.CtrlGetLink)
	aCtx.App.Post("/*", aCtx.CtrlPostLink)
}

func NewClient(ctx context.Context, c config.Config) *AppContext {
	aCtx := &AppContext{
		App:   fiber.New(),
		LinkS: usecase.NewLinkService(ctx, c),
	}
	InitControllers(aCtx)
	return aCtx
}
