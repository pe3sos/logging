package formatter

import (
	"strings"

	"github.com/zippunov/logging/def"
)

func Compose(fms ...def.Formatter) def.Formatter {
	l := len(fms)
	reverse := make([]def.Formatter, len(fms))
	for idx, el := range fms {
		reverse[l-idx-1] = el
	}
	return &composableFormatter{fms, reverse}
}

type composableFormatter struct {
	straight []def.Formatter
	reverse  []def.Formatter
}

func (f *composableFormatter) GetPrefix(lvl def.Level) string {
	var sb strings.Builder
	l := len(f.straight)
	for idx, formatter := range f.straight {
		sb.WriteString(formatter.GetPrefix(lvl))
		if idx < l-1 {
			sb.WriteRune(' ')
		}
	}
	return sb.String()
}

// GetSuffix returns ""
func (f *composableFormatter) GetSuffix(lvl def.Level) string {
	var sb strings.Builder
	for _, formatter := range f.reverse {
		sb.WriteString(formatter.GetSuffix(lvl))
	}
	return sb.String()
}

// Format adds filename and line number before the log message
func (f *composableFormatter) Format(lvl def.Level, format string, values ...interface{}) (string, []interface{}) {
	for _, formatter := range f.reverse {
		format, values = formatter.Format(lvl, format, values...)
	}
	return format, values
}
