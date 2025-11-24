package numbers

import (
	"fmt"
	"net/http"

	"github.com/avraam311/profiled-number-adder/internal/api/http/responses"
	"github.com/avraam311/profiled-number-adder/internal/infra/logger"
	"github.com/avraam311/profiled-number-adder/internal/models/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddUp(c *gin.Context) {
	var nums dto.Numbers

	// Use ShouldBindBodyWith with JSON validator to avoid repeated unmarshaling on middleware if applicable
	if err := c.ShouldBindJSON(&nums); err != nil {
		// Log only on error to reduce logging overhead
		logger.Logger.Error().Err(err).Msg("failed to decode or validate request body")
		responses.ResponseError(c, responses.ErrInvalidJSON, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	sum, err := h.service.AddUp(c.Request.Context(), &nums)
	if err != nil {
		// Log errors with context only
		logger.Logger.Error().Err(err).Interface("nums", nums).Msg("failed to add up numbers")
		responses.ResponseError(c, responses.ErrInternalServer, "internal server error", http.StatusInternalServerError)
		return
	}

	responses.ResponseCreated(c, sum)
}
