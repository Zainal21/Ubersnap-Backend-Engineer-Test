package middleware

import (
	"encoding/base64"
	"strings"

	"github.com/Zainal21/Ubersnap-backend-test/app/appctx"
	"github.com/Zainal21/Ubersnap-backend-test/pkg/config"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (a *AuthMiddleware) Authenticate(xCtx *fiber.Ctx, conf *config.Config) appctx.DynamicResponse {
	auth := xCtx.GetReqHeaders()["Authorization"]

	if len(auth) == 0 {
		return *appctx.NewResponse().WithCode(fiber.StatusUnauthorized).WithMessage("Unauthorized")
	}

	decodeString, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic "))
	if err != nil {
		return *appctx.NewResponse().WithCode(fiber.StatusUnauthorized).WithMessage("Unauthorized")
	}

	resultAuth := strings.Split(string(decodeString), ":")

	if resultAuth[0] == "username" && resultAuth[1] == "password" {
		return *appctx.NewResponse().WithCode(fiber.StatusOK)
	}

	return *appctx.NewResponse().WithCode(fiber.StatusUnauthorized).WithMessage("Unauthorized")
}
