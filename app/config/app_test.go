package config

import (
	"testing"
)

func TestConfiguration(t *testing.T) {
	conf := Configuration()

	if &conf == nil {
		t.Error("Failed to parse config.yml file")
	}
}

func BenchmarkConfiguration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Configuration()
	}
}
