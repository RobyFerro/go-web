package config

import (
	"testing"
)

func TestConfigurationWeb(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "Get routing",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConfigurationWeb(); &got == nil {
				t.Errorf("ConfigurationWeb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkConfigurationWeb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ConfigurationWeb()
	}
}
