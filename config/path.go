package config

import "path/filepath"

var (
	_, b, _, _ = runtime.Caller(0)
	bPath      = filepath.Dir(b)
)

func GetFilePath(path string) string {
	return filepath.Join(filepath.Join(bPath, "../../"), path)
}

func GetBasePath() string {
	return bPath
}
