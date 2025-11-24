package benchmarks

import (
	"context"
	"testing"

	"github.com/avraam311/profiled-number-adder/internal/models/dto"
	"github.com/avraam311/profiled-number-adder/internal/service/numbers"
)

func BenchmarkAddUp(b *testing.B) {
	svc := numbers.New()
	nums := &dto.Numbers{Num1: 100, Num2: 200}

	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		_, err := svc.AddUp(ctx, nums)
		if err != nil {
			b.Fatalf("AddUp returned error: %v", err)
		}
	}
}
