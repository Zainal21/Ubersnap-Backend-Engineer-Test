package controller

import (
	"github.com/Zainal21/Ubersnap-backend-test/app/appctx"
	"github.com/Zainal21/Ubersnap-backend-test/app/controller/contract"
	"github.com/gofiber/fiber/v2"
)

type getHealth struct {
}

func (g *getHealth) Serve(xCtx appctx.Data) appctx.Response {
	// Ping Endpoint
	return *appctx.NewResponse().WithCode(fiber.StatusOK).WithMessage("OK").WithData(struct {
		Message string `json:"message"`
	}{
		Message: "healhty!",
	})
}

func NewGetHealth() contract.Controller {
	return &getHealth{}
}
