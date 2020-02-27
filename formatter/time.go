package formatter

import (
	"strings"
	"time"

	"github.com/zippunov/logging/def"
)

func Time() def.Formatter {
	return &timeFormatter{}
}

type timeFormatter struct {
}

func (f *timeFormatter) GetPrefix(lvl def.Level) string {
	t := time.Now()
	var sb strings.Builder
	year, month, day := t.Date()
	sb.Write(itoa(year, 4))
	sb.WriteRune('/')
	sb.Write(itoa(int(month), 2))
	sb.WriteRune('/')
	sb.Write(itoa(day, 2))
	sb.WriteRune(' ')
	hour, min, sec := t.Clock()
	sb.Write(itoa(hour, 2))
	sb.WriteRune(':')
	sb.Write(itoa(min, 2))
	sb.WriteRune(':')
	sb.Write(itoa(sec, 2))
	return sb.String()
}

func (f *timeFormatter) GetSuffix(lvl def.Level) string {
	return ""
}

// Format modifies format string and format params list
func (f *timeFormatter) Format(lvl def.Level, format string, values ...interface{}) (string, []interface{}) {
	return format, values
}
