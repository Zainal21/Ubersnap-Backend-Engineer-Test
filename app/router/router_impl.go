package router

import (
	"github.com/Zainal21/Ubersnap-backend-test/app/appctx"
	"github.com/Zainal21/Ubersnap-backend-test/app/controller"
	"github.com/Zainal21/Ubersnap-backend-test/app/controller/contract"
	"github.com/Zainal21/Ubersnap-backend-test/app/controller/image_processing"
	"github.com/Zainal21/Ubersnap-backend-test/app/handler"
	"github.com/Zainal21/Ubersnap-backend-test/app/middleware"
	"github.com/Zainal21/Ubersnap-backend-test/app/service"
	"github.com/Zainal21/Ubersnap-backend-test/pkg/config"

	"github.com/gofiber/fiber/v2"
)

type router struct {
	cfg   *config.Config
	fiber fiber.Router
}

func (rtr *router) handle(hfn httpHandlerFunc, svc contract.Controller, mdws ...middleware.MiddlewareFunc) fiber.Handler {
	return func(xCtx *fiber.Ctx) error {

		//check registered middleware functions
		if rm := middleware.FilterFunc(rtr.cfg, xCtx, mdws); rm.Code != fiber.StatusOK {
			// return response base on middleware
			res := *appctx.NewResponse().
				WithCode(rm.Code).
				WithError(rm.Errors).
				WithMessage(rm.Message)
			return rtr.response(xCtx, res)
		}

		//send to controller
		resp := hfn(xCtx, svc, rtr.cfg)
		return rtr.response(xCtx, resp)
	}
}

func (rtr *router) customHandle(hfn httpApiCustomHandlerFunc, svc contract.CustomResponseController, mdws ...middleware.MiddlewareApiCoreDukcapilFunc) fiber.Handler {
	return func(xCtx *fiber.Ctx) error {
		// Check registered middleware functions
		rm := middleware.FilterCoreCustomFunc(rtr.cfg, xCtx, mdws)

		switch v := rm.(type) {
		case appctx.Response:
			if v.Code != fiber.StatusOK {
				// Return response based on middleware
				return rtr.customReponse(xCtx, v.Code, map[string]interface{}{
					"status": false,
				})
			}
		case map[string]interface{}:
			// Handle other types,  map[string]interface{}
			code, _ := v["code"].(int)
			if code != fiber.StatusOK {
				message, _ := v["message"].(string)

				return rtr.customReponse(xCtx, code, map[string]interface{}{
					"status":  false,
					"code":    code,
					"message": message,
				})
			}
		default:
			return xCtx.JSON(v)
		}

		// Send to controller
		resp := hfn(xCtx, svc, rtr.cfg)
		// Check the type of the response and return accordingly
		switch respObj := resp.(type) {
		case appctx.Response:
			return rtr.customReponse(xCtx, respObj.Code, respObj)
		case []byte:
			return xCtx.JSON(respObj)
		case string:
			return xCtx.SendString(respObj)
		default:
			return rtr.customReponse(xCtx, fiber.StatusOK, respObj)
		}
	}
}

func (rtr *router) customReponse(fiberCtx *fiber.Ctx, statusCode int, response interface{}) error {
	fiberCtx.Status(statusCode)
	return fiberCtx.JSON(response)
}

func (rtr *router) response(fiberCtx *fiber.Ctx, resp appctx.Response) error {
	fiberCtx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return fiberCtx.Status(resp.Code).Send(resp.Byte())
}

func (rtr *router) Route() {
	//init db
	// db := bootstrap.RegistryDatabase(rtr.cfg)

	//define repositories
	// userRepo := repositories.NewUserRepositoryImpl(db)
	// tokenRepo := repositories.NewPersonalToken(db, &sanctum.Token{
	// 	Crypto: &cryptoservice.Crypto{},
	// }, userRepo)

	//define services
	imageService := service.NewImageProcessingServiceImpl()

	//define middleware
	basicMiddleware := middleware.NewAuthMiddleware()

	//define controller
	convert := image_processing.NewConvert(imageService)
	resize := image_processing.NewResize(imageService)
	compress := image_processing.NewCompress(imageService)

	health := controller.NewGetHealth()
	privateV1 := rtr.fiber.Group("/api/private/v1")

	rtr.fiber.Get("/up", rtr.handle(
		handler.HttpRequest,
		health,
	))

	privateV1.Post("/convert", rtr.customHandle(
		handler.HttpApiCustomRequest,
		convert,
		func(xCtx *fiber.Ctx, conf *config.Config) any {
			return basicMiddleware
		},
	))

	privateV1.Post("/resize", rtr.customHandle(
		handler.HttpApiCustomRequest,
		resize,
		func(xCtx *fiber.Ctx, conf *config.Config) any {
			return basicMiddleware
		},
	))

	privateV1.Post("/compress", rtr.customHandle(
		handler.HttpApiCustomRequest,
		compress,
		func(xCtx *fiber.Ctx, conf *config.Config) any {
			return basicMiddleware
		},
	))
}

func NewRouter(cfg *config.Config, fiber fiber.Router) Router {
	return &router{
		cfg:   cfg,
		fiber: fiber,
	}
}
