package numbers

import (
	"context"

	"github.com/avraam311/profiled-number-adder/internal/models/dto"
)

type Service interface {
	AddUp(context.Context, *dto.Numbers) (int, error)
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{
		service: service,
	}
}
