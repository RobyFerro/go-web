package config

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	bPath      = filepath.Dir(b)
)

// Returns the absolute path of the selected file/folder.
// The basic path is Go-Web main folder.
// Example: GetDynamicPath("storage/certs/tls.key")
func GetDynamicPath(path string) string {
	return filepath.Join(filepath.Join(bPath, "../../"), path)
}

// Return the basic project path.
// Deprecated: obsolete method, use the GetDynamicPath instead.
func GetBasePath() string {
	return bPath
}
