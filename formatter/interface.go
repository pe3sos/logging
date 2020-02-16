package formatter

import "github.com/zippunov/logging/def"

// Formatter interface
type Formatter interface {
	GetPrefix(lvl def.Level) string
	Format(lvl def.Level, format string, values ...interface{}) (string, []interface{})
	GetSuffix(lvl def.Level) string
}
