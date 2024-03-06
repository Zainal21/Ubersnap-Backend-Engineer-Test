package router

import (
	"github.com/Zainal21/Ubersnap-backend-test/app/appctx"
	"github.com/Zainal21/Ubersnap-backend-test/app/controller/contract"
	"github.com/Zainal21/Ubersnap-backend-test/pkg/config"
	"github.com/gofiber/fiber/v2"
)

type httpHandlerFunc func(xCtx *fiber.Ctx, svc contract.Controller, conf *config.Config) appctx.Response

type httpStreamHandlerFunc func(xCtx *fiber.Ctx, svc contract.StreamController, conf *config.Config) appctx.DynamicResponse

type httpApiCustomHandlerFunc func(xCtx *fiber.Ctx, svc contract.CustomResponseController, conf *config.Config) appctx.DynamicResponse

type Router interface {
	Route()
}
