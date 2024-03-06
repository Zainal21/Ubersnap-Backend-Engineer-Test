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

type resizeImpl struct {
	service service.ImageProcessingService
}

func (g *resizeImpl) Serve(xCtx appctx.Data) appctx.DynamicResponse {
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

	// width resize
	widthResize, err := helpers.GetParameter(formData, "width")
	if err != nil {
		return helpers.HandleErrorResponse(err.Error(), fiber.StatusBadRequest)
	}

	widthResizeFinal, err := strconv.Atoi(widthResize)
	if err != nil {
		return helpers.HandleErrorResponse("Invalid 'quality' parameter: "+err.Error(), fiber.StatusBadRequest)
	}

	heightResize, ok := formData.Value["width"]
	if !ok {
		return helpers.HandleErrorResponse("Invalid Parameter", fiber.StatusBadRequest)
	}

	heightResizeFinal, err := strconv.Atoi(heightResize[0])
	if err != nil {
		return helpers.HandleErrorResponse("Invalid Parameter", fiber.StatusBadRequest)
	}

	return g.service.ResizeImage(xCtx.FiberCtx.Context(), files, widthResizeFinal, heightResizeFinal)
}

func NewResize(
	service service.ImageProcessingService,
) contract.StreamController {
	return &resizeImpl{
		service: service,
	}
}
