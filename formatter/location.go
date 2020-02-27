package formatter

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/zippunov/logging/def"
)

func Location(depth int) def.Formatter {
	return &locationFormatter{Depth: depth}
}

// LocationFormatter adds filename and line number before the log message
type locationFormatter struct {
	Depth int
}

// GetPrefix returns ""
func (f *locationFormatter) GetPrefix(lvl def.Level) string {
	return f.header()
}

// GetSuffix returns ""
func (f *locationFormatter) GetSuffix(lvl def.Level) string {
	return ""
}

// Format modifies format string and format params list
func (f *locationFormatter) Format(lvl def.Level, format string, values ...interface{}) (string, []interface{}) {
	return format, values
}

// Returns header including filename and line number
func (f *locationFormatter) header() string {
	_, fn, line, ok := runtime.Caller(f.Depth)
	if !ok {
		fn = "???"
		line = 1
	}

	return fmt.Sprintf("%s:%d", filepath.Base(fn), line)
}
