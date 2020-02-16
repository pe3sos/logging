package formatter

import "github.com/zippunov/logging/def"

func Default() Formatter {
	return &defaultFormatter{}
}

type defaultFormatter struct {
}

// GetPrefix returns ""
func (f *defaultFormatter) GetPrefix(lvl def.Level) string {
	return ""
}

// GetSuffix returns ""
func (f *defaultFormatter) GetSuffix(lvl def.Level) string {
	return ""
}

// Format adds filename and line number before the log message
func (f *defaultFormatter) Format(lvl def.Level, format string, values ...interface{}) (string, []interface{}) {
	return format, values
}
