package numbers

import (
	"context"

	"github.com/avraam311/profiled-number-adder/internal/models/dto"
)

func (s *Service) AddUp(ctx context.Context, nums *dto.Numbers) (int, error) {
	num1 := nums.Num1
	num2 := nums.Num2
	sum := num1 + num2

	return sum, nil
}
