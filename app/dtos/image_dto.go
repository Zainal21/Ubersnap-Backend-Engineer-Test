package dtos

type ConvertDto struct {
	Image string `json:"image" validate:"required"`
}
