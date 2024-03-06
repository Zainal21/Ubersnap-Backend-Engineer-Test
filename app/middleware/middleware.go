package middleware

import (
	"github.com/Zainal21/Ubersnap-backend-test/app/appctx"
	"github.com/Zainal21/Ubersnap-backend-test/pkg/config"
	"github.com/gofiber/fiber/v2"
)

type MiddlewareFunc func(xCtx *fiber.Ctx, conf *config.Config) appctx.Response

type MiddlewareStreamFunc func(xCtx *fiber.Ctx, conf *config.Config) any // handle for custom response

type MiddlewareApiCoreDukcapilFunc func(xCtx *fiber.Ctx, conf *config.Config) any // handle for custom response

// filter func for custom response
func FilterCoreCustomFunc(conf *config.Config, xCtx *fiber.Ctx, mfs []MiddlewareApiCoreDukcapilFunc) interface{} {
	response := map[string]interface{}{
		"code": fiber.StatusOK,
	}
	for _, mf := range mfs {
		if result := mf(xCtx, conf); result != nil {
			resultMap := handleMiddlewareResult(result)
			if resultMap != nil {
				if code, codeOk := resultMap["code"].(int); codeOk && code != fiber.StatusOK {
					return resultMap
				}
				if message, messageOk := resultMap["message"].(string); messageOk {
					response["message"] = message
				}
			}
		}
	}

	return response
}

// FilterFunc is a iterator resolver in each middleware registered
func FilterFunc(conf *config.Config, xCtx *fiber.Ctx, mfs []MiddlewareFunc) appctx.Response {
	// Initiate postive case
	var response = appctx.Response{Code: fiber.StatusOK}
	for _, mf := range mfs {
		if response = mf(xCtx, conf); response.Code != fiber.StatusOK {
			return response
		}
	}

	return response
}

func handleMiddlewareResult(result interface{}) map[string]interface{} {
	if result != nil {
		switch val := result.(type) {
		case appctx.Response:
			return map[string]interface{}{
				"code":    val.Code,
				"message": val.Message,
			}
		case map[string]interface{}:
			if message, messageOk := val["message"].(string); messageOk {
				return map[string]interface{}{
					"code":    val["code"],
					"message": message,
				}
			}
			if code, codeOk := val["code"].(int); codeOk {
				return map[string]interface{}{
					"code": code,
				}
			}
		}
	}
	return nil
}
