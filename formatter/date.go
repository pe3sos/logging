package formatter

import (
	"strings"
	"time"

	"github.com/zippunov/logging/def"
)

func Date() def.Formatter {
	return &dateFormatter{}
}

type dateFormatter struct {
}

func (f *dateFormatter) GetPrefix(lvl def.Level) string {
	t := time.Now()
	var sb strings.Builder
	year, month, day := t.Date()
	sb.Write(itoa(year, 4))
	sb.WriteRune('/')
	sb.Write(itoa(int(month), 2))
	sb.WriteRune('/')
	sb.Write(itoa(day, 2))
	return sb.String()
}

func (f *dateFormatter) GetSuffix(lvl def.Level) string {
	return ""
}

func (f *dateFormatter) Format(lvl def.Level, format string, values ...interface{}) (string, []interface{}) {
	return format, values
}
