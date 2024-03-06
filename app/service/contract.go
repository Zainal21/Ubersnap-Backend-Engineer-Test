package service

import (
	"context"
	"mime/multipart"

	"github.com/Zainal21/Ubersnap-backend-test/app/appctx"
)

type UserService interface {
}

type ImageProcessingService interface {
	// convert png to jpeg
	ConvertPNGtoJPEG(ctx context.Context, files []*multipart.FileHeader) appctx.Response
	// resize image with custom width & height
	ResizeImage(ctx context.Context, files []*multipart.FileHeader, width, height int) appctx.Response
	// compress image with quality parameter
	CompressImage(ctx context.Context, files []*multipart.FileHeader, quality int) appctx.Response
}
