package benchmarks

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	handlerNumbers "github.com/avraam311/profiled-number-adder/internal/api/http/handlers/numbers"
	svcNumbers "github.com/avraam311/profiled-number-adder/internal/service/numbers"
	"github.com/gin-gonic/gin"
)

func BenchmarkAddUpHandler(b *testing.B) {
	gin.SetMode(gin.TestMode)

	svc := svcNumbers.New()
	handler := handlerNumbers.New(svc)

	for i := 0; i < b.N; i++ {
		numsJSON := `{"num_1":100,"num_2":200}`
		req, _ := http.NewRequest(http.MethodPost, "/add-up", bytes.NewBufferString(numsJSON))
		req = req.WithContext(context.Background())
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req

		handler.AddUp(c)

		if w.Code != http.StatusCreated {
			b.Fatalf("expected status %d but got %d", http.StatusCreated, w.Code)
		}
	}
}
