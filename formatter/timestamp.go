package formatter

import (
	"strings"
	"time"

	"github.com/zippunov/logging/def"
)

const nanosInMillis int = 1e3

// Cheap integer to fixed-width decimal ASCII. Give a negative width to avoid zero-padding.
func itoa(i, wid int) []byte {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	return b[bp:]
}

func Timestamp() def.Formatter {
	return &timestampFormatter{}
}

type timestampFormatter struct {
}

func (f *timestampFormatter) GetPrefix(lvl def.Level) string {
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
	sb.WriteRune('.')
	sb.Write(itoa(t.Nanosecond()/nanosInMillis, 6))
	return sb.String()
}

func (f *timestampFormatter) GetSuffix(lvl def.Level) string {
	return ""
}

// Format modifies format string and format params list
func (f *timestampFormatter) Format(lvl def.Level, format string, values ...interface{}) (string, []interface{}) {
	return format, values
}
