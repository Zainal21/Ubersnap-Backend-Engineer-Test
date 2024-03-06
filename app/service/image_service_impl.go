package service

import (
	"context"
	"image"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/Zainal21/Ubersnap-backend-test/app/appctx"
	"github.com/Zainal21/Ubersnap-backend-test/app/consts"
	"github.com/Zainal21/Ubersnap-backend-test/app/helpers"
	"github.com/gofiber/fiber/v2"
	"gocv.io/x/gocv"
)

type ImageProcessingServiceImpl struct{}

// ResizeImage implements ImageProcessingService.
func (i *ImageProcessingServiceImpl) ResizeImage(ctx context.Context, files []*multipart.FileHeader, width, height int) appctx.Response {
	convertedFiles := make([]string, 0)
	for _, file := range files {
		filename := file.Filename
		fileBytes, err := OpenFile(file)
		if err != nil {
			return helpers.HandleErrorResponse("Failed to open or read the uploaded file", fiber.StatusBadRequest)
		}

		imageMat, err := gocv.IMDecode(fileBytes, gocv.IMReadColor)
		if err != nil {
			return helpers.HandleErrorResponse("Error decoding PNG image with Gocv", fiber.StatusBadRequest)
		}
		defer imageMat.Close()

		resizeMatImg := gocv.NewMat()
		gocv.CvtColor(imageMat, &resizeMatImg, gocv.ColorBGRToGray)
		gocv.Resize(imageMat, &resizeMatImg, image.Pt(width, height), 0, 0, gocv.InterpolationDefault)
		defer resizeMatImg.Close()

		outputPath := consts.ResizeDirectory + "/" + filename
		if _, err := os.Stat(outputPath); err == nil {
			if err := os.Remove(outputPath); err != nil {
				return helpers.HandleErrorResponse("Error deleting existing file", fiber.StatusInternalServerError)
			}
		}
		if ok := gocv.IMWrite(outputPath, resizeMatImg); !ok {
			return helpers.HandleErrorResponse("Error writing JPEG image with Gocv", fiber.StatusBadRequest)
		}

		encodedImage, err := helpers.ConvertImageToBase64(outputPath)
		if err != nil {
			return helpers.HandleErrorResponse("Error converting image to base64", fiber.StatusInternalServerError)
		}

		convertedFiles = append(convertedFiles, encodedImage)
	}

	return *appctx.NewResponse().WithCode(fiber.StatusOK).WithMessage("SUCCESS").WithData(fiber.Map{"message": "File(s) resize and saved successfully", "image": convertedFiles[0]})
}

// OpenFile opens the uploaded file and returns its content as []byte.
func OpenFile(file *multipart.FileHeader) ([]byte, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

// WriteFile writes the given data to the specified file path.
func (i *ImageProcessingServiceImpl) WriteFile(filePath string, data []byte) error {
	return ioutil.WriteFile(filePath, data, 0644)
}

// CompressImage implements ImageProcessingService.
func (i *ImageProcessingServiceImpl) CompressImage(ctx context.Context, files []*multipart.FileHeader, quality int) appctx.Response {
	convertedFiles := make([]string, 0)

	for _, file := range files {
		filename := file.Filename

		fileBytes, err := OpenFile(file)
		if err != nil {
			return helpers.HandleErrorResponse("Failed to open or read the uploaded file ", fiber.StatusBadRequest)
		}

		imageMat, err := gocv.IMDecode(fileBytes, gocv.IMReadColor)
		if err != nil {
			return helpers.HandleErrorResponse("Error decoding PNG image with Gocv", fiber.StatusBadRequest)
		}
		defer imageMat.Close()

		compressing := []int{gocv.IMWriteJpegQuality, quality}
		encodedImg, err := gocv.IMEncodeWithParams(".jpg", imageMat, compressing)
		if err != nil {
			return helpers.HandleErrorResponse("Error encoding image", fiber.StatusBadRequest)
		}

		outputPath := consts.CompressDirectory + "/" + filename
		if _, err := os.Stat(outputPath); err == nil {
			if err := os.Remove(outputPath); err != nil {
				return helpers.HandleErrorResponse("Error deleting existing file", fiber.StatusInternalServerError)
			}
		}
		err = i.WriteFile(outputPath, encodedImg.GetBytes())
		if err != nil {
			return helpers.HandleErrorResponse("Error writing to file", fiber.StatusInternalServerError)
		}

		encodedImage, err := helpers.ConvertImageToBase64(outputPath)
		if err != nil {
			return helpers.HandleErrorResponse("Error converting image to base64", fiber.StatusInternalServerError)
		}

		convertedFiles = append(convertedFiles, encodedImage)
	}

	return *appctx.NewResponse().WithCode(fiber.StatusOK).WithMessage("SUCCESS").WithData(fiber.Map{"message": "File(s) compressed and saved successfully", "image": convertedFiles[0]})
}

// ConvertPNGtoJPEG implements ImageProcessingService.
func (i *ImageProcessingServiceImpl) ConvertPNGtoJPEG(ctx context.Context, files []*multipart.FileHeader) appctx.Response {
	convertedFiles := make([]string, 0)
	for _, file := range files {
		filename := file.Filename

		fileBytes, err := OpenFile(file)
		if err != nil {
			return helpers.HandleErrorResponse("Failed to open or read the uploaded file", fiber.StatusBadRequest)
		}

		imageMat, err := gocv.IMDecode(fileBytes, gocv.IMReadColor)
		if err != nil {
			return helpers.HandleErrorResponse("Error decoding PNG image with Gocv", fiber.StatusBadRequest)
		}

		outputPath := consts.ConvertDirectory + "/" + strings.TrimSuffix(filename, filepath.Ext(filename)) + ".jpeg"
		if _, err := os.Stat(outputPath); err == nil {
			if err := os.Remove(outputPath); err != nil {
				return helpers.HandleErrorResponse("Error deleting existing file", fiber.StatusInternalServerError)
			}
		}

		if ok := gocv.IMWrite(outputPath, imageMat); !ok {
			return helpers.HandleErrorResponse("Error writing PNG image with Gocv", fiber.StatusBadRequest)
		}

		encodedImage, err := helpers.ConvertImageToBase64(outputPath)
		if err != nil {
			return helpers.HandleErrorResponse("Error converting image to base64", fiber.StatusInternalServerError)
		}

		convertedFiles = append(convertedFiles, encodedImage)
	}

	return *appctx.NewResponse().WithCode(fiber.StatusOK).WithMessage("SUCCESS").WithData(fiber.Map{"message": "File(s) converted and saved successfully", "image": convertedFiles[0]})
}

func NewImageProcessingServiceImpl() ImageProcessingService {
	return &ImageProcessingServiceImpl{}
}
