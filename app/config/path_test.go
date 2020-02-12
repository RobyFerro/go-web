package config

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestGetBasePath(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "true",
			want: filepath.Join(bPath, "../.."),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBasePath(); got != tt.want {
				t.Errorf("GetBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGetBasePath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetBasePath()
	}
}

func TestGetDynamicPath(t *testing.T) {
	type args struct {
		path string
	}

	a := args{path: "config.yml"}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Get config file",
			args: a,
			want: fmt.Sprintf("%s/config.yml", filepath.Join(bPath, "../..")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDynamicPath(tt.args.path); got != tt.want {
				t.Errorf("GetDynamicPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGetDynamicPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetDynamicPath("/database")
	}
}
