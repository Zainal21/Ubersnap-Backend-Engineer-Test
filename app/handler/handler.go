package handler

import (
	"github.com/Zainal21/Ubersnap-backend-test/app/appctx"
	"github.com/Zainal21/Ubersnap-backend-test/app/controller/contract"
	"github.com/Zainal21/Ubersnap-backend-test/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func HttpRequest(xCtx *fiber.Ctx, svc contract.Controller, conf *config.Config) appctx.Response {
	data := appctx.Data{
		FiberCtx: xCtx,
		Cfg:      conf,
	}
	return svc.Serve(data)
}

func HttpApiCustomRequest(xCtx *fiber.Ctx, svc contract.CustomResponseController, conf *config.Config) appctx.DynamicResponse {
	data := appctx.Data{
		FiberCtx: xCtx,
		Cfg:      conf,
	}
	return svc.Serve(data)
}

func HttpApiStreamRequest(xCtx *fiber.Ctx, svc contract.StreamController, conf *config.Config) appctx.DynamicResponse {
	data := appctx.Data{
		FiberCtx: xCtx,
		Cfg:      conf,
	}
	return svc.Serve(data)
}
