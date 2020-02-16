package logging

import (
	"io"
	"os"

	"github.com/zippunov/logging/def"
	fmr "github.com/zippunov/logging/formatter"
)

type Logger struct {
	maxLvl  def.Level
	DEBUG   LoggerInterface
	INFO    LoggerInterface
	WARNING LoggerInterface
	ERROR   LoggerInterface
	FATAL   LoggerInterface
}

func (l *Logger) GetMaxLevel() def.Level {
	return l.maxLvl
}

func (l *Logger) SetMaxLevel(lvl def.Level) {
	l.maxLvl = lvl
}

func New(out, errOut io.Writer, f fmr.Formatter) *Logger {

	if out == nil {
		out = os.Stdout
	}
	if errOut == nil {
		errOut = os.Stderr
	}
	if f == nil {
		f = fmr.Default()
	}
	l := &Logger{maxLvl: def.DEBUG}
	l.DEBUG = Appender(out, def.DEBUG, f, l)
	l.INFO = Appender(out, def.INFO, f, l)
	l.WARNING = Appender(out, def.WARNING, f, l)
	l.ERROR = Appender(errOut, def.ERROR, f, l)
	l.FATAL = Appender(errOut, def.ERROR, f, l)
	return l
}
