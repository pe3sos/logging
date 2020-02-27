package formatter

import (
	"fmt"

	"github.com/zippunov/logging/def"
)

const (
	// For coloring
	resetSeq  = "\033[0m"
	colourSeq = "\033[0;%dm"
)

// Color map
var color = map[def.Level]string{
	def.Levels.DEBUG:   fmt.Sprintf(colourSeq, 37), // grey
	def.Levels.INFO:    fmt.Sprintf(colourSeq, 94), // blue
	def.Levels.WARNING: fmt.Sprintf(colourSeq, 95), // pink
	def.Levels.ERROR:   fmt.Sprintf(colourSeq, 91), // red
	def.Levels.FATAL:   fmt.Sprintf(colourSeq, 91), // red
}

func Coloured() def.Formatter {
	return &coloredFormatter{}
}

// ColoredFormatter colors log messages with ASCI escape codes
// and adds filename and line number before the log message
// See https://en.wikipedia.org/wiki/ANSI_escape_code
type coloredFormatter struct {
}

// GetPrefix returns color escape code
func (f *coloredFormatter) GetPrefix(lvl def.Level) string {
	return color[lvl]
}

// GetSuffix returns reset sequence code
func (f *coloredFormatter) GetSuffix(lvl def.Level) string {
	return resetSeq
}

// Format adds filename and line number before the log message
func (f *coloredFormatter) Format(lvl def.Level, format string, values ...interface{}) (string, []interface{}) {
	return format, values
}
