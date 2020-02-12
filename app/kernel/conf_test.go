package kernel

import (
	"testing"
)

func TestWebRouter(t *testing.T) {
	if WebRouter() == nil {
		t.Error("Failed to parse router")
	}
}

func BenchmarkWebRouter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = WebRouter()
	}
}
