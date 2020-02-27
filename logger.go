package logging

import (
	"io"
	"os"

	"github.com/zippunov/logging/def"
	fmr "github.com/zippunov/logging/formatter"
)

type Logger struct {
	maxLvl  def.Level
	DEBUG   def.LoggerInterface
	INFO    def.LoggerInterface
	WARNING def.LoggerInterface
	ERROR   def.LoggerInterface
	FATAL   def.LoggerInterface
}

func (l *Logger) GetMaxLevel() def.Level {
	return l.maxLvl
}

func (l *Logger) SetMaxLevel(lvl def.Level) {
	l.maxLvl = lvl
}

func New(out, errOut io.Writer, f def.Formatter) *Logger {
	if out == nil {
		out = os.Stdout
	}
	if errOut == nil {
		errOut = os.Stderr
	}
	if f == nil {
		f = fmr.Default()
	}
	l := &Logger{maxLvl: def.Levels.DEBUG}
	l.DEBUG = Appender(out, def.Levels.DEBUG, f, l)
	l.INFO = Appender(out, def.Levels.INFO, f, l)
	l.WARNING = Appender(out, def.Levels.WARNING, f, l)
	l.ERROR = Appender(errOut, def.Levels.ERROR, f, l)
	l.FATAL = Appender(errOut, def.Levels.FATAL, f, l)
	return l
}
