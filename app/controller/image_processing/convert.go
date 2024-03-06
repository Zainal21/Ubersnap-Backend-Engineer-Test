package image_processing

import (
	"os"

	"github.com/Zainal21/Ubersnap-backend-test/app/appctx"
	"github.com/Zainal21/Ubersnap-backend-test/app/consts"
	"github.com/Zainal21/Ubersnap-backend-test/app/controller/contract"
	"github.com/Zainal21/Ubersnap-backend-test/app/helpers"
	"github.com/Zainal21/Ubersnap-backend-test/app/service"
	"github.com/gofiber/fiber/v2"
)

type convertImpl struct {
	service service.ImageProcessingService
}

func (c *convertImpl) Serve(xCtx appctx.Data) appctx.DynamicResponse {
	formData, err := xCtx.FiberCtx.MultipartForm()
	if err != nil {
		return helpers.HandleErrorResponse("Failed to parse form data", fiber.StatusBadRequest)
	}
	// crearte upload directory
	if err := os.MkdirAll(consts.UploadDirectory, os.ModePerm); err != nil {
		return helpers.HandleErrorResponse("Failed to create upload directory", fiber.StatusInternalServerError)
	}

	files, ok := formData.File["image"]

	if !ok {
		return helpers.HandleErrorResponse("Invalid Parameter", fiber.StatusUnprocessableEntity)
	}

	return c.service.ConvertPNGtoJPEG(xCtx.FiberCtx.Context(), files)
}

func NewConvert(
	service service.ImageProcessingService,
) contract.CustomResponseController {
	return &convertImpl{
		service: service,
	}
}
