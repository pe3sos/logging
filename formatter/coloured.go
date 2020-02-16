package formatter

import (
	"fmt"

	"github.com/zippunov/logging/def"
)

const (
	// For colouring
	resetSeq  = "\033[0m"
	colourSeq = "\033[0;%dm"
)

// Colour map
var colour = map[def.Level]string{
	def.DEBUG:   fmt.Sprintf(colourSeq, 37), // grey
	def.INFO:    fmt.Sprintf(colourSeq, 94), // blue
	def.WARNING: fmt.Sprintf(colourSeq, 95), // pink
	def.ERROR:   fmt.Sprintf(colourSeq, 91), // red
	def.FATAL:   fmt.Sprintf(colourSeq, 91), // red
}

func Coloured() Formatter {
	return &colouredFormatter{}
}

// ColouredFormatter colours log messages with ASCI escape codes
// and adds filename and line number before the log message
// See https://en.wikipedia.org/wiki/ANSI_escape_code
type colouredFormatter struct {
}

// GetPrefix returns colour escape code
func (f *colouredFormatter) GetPrefix(lvl def.Level) string {
	return colour[lvl]
}

// GetSuffix returns reset sequence code
func (f *colouredFormatter) GetSuffix(lvl def.Level) string {
	return resetSeq
}

// Format adds filename and line number before the log message
func (f *colouredFormatter) Format(lvl def.Level, format string, values ...interface{}) (string, []interface{}) {
	return format, values
}
