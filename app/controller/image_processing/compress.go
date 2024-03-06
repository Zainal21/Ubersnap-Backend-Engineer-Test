package image_processing

import (
	"os"
	"strconv"

	"github.com/Zainal21/Ubersnap-backend-test/app/appctx"
	"github.com/Zainal21/Ubersnap-backend-test/app/consts"
	"github.com/Zainal21/Ubersnap-backend-test/app/controller/contract"
	"github.com/Zainal21/Ubersnap-backend-test/app/helpers"
	"github.com/Zainal21/Ubersnap-backend-test/app/service"
	"github.com/gofiber/fiber/v2"
)

type compressImpl struct {
	service service.ImageProcessingService
}

func (g *compressImpl) Serve(xCtx appctx.Data) appctx.DynamicResponse {
	formData, err := xCtx.FiberCtx.MultipartForm()
	if err != nil {
		return helpers.HandleErrorResponse("Failed to parse form data", fiber.StatusBadRequest)
	}
	// crearte upload directory
	if err := os.MkdirAll(consts.UploadDirectory, os.ModePerm); err != nil {
		return helpers.HandleErrorResponse("Failed to create upload directory", fiber.StatusBadRequest)
	}

	files, ok := formData.File["image"]
	if !ok {
		return helpers.HandleErrorResponse("Invalid Parameter", fiber.StatusBadRequest)
	}

	quality, err := helpers.GetParameter(formData, "quality")
	if err != nil {
		return helpers.HandleErrorResponse(err.Error(), fiber.StatusBadRequest)
	}

	imageQualityFinal, err := strconv.Atoi(quality)
	if err != nil {
		return helpers.HandleErrorResponse("Invalid 'quality' parameter: "+err.Error(), fiber.StatusBadRequest)
	}

	return g.service.CompressImage(xCtx.FiberCtx.Context(), files, imageQualityFinal)
}

func NewCompress(
	service service.ImageProcessingService,
) contract.StreamController {
	return &compressImpl{
		service: service,
	}
}
