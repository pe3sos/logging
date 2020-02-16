package logging

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/zippunov/logging/def"
	fmr "github.com/zippunov/logging/formatter"
)

type appender struct {
	registry  *Logger
	lvl       def.Level
	formatter fmr.Formatter
	out       io.Writer  // destination for output
	buf       []byte     // for accumulating text to write
	mu        sync.Mutex // ensures atomic writes; protects the following fields
}

// New creates a new Appender. The out variable sets the
// destination to which log data will be written.
// The prefix appears at the beginning of each generated log line.
// The flag argument defines the logging properties.
func Appender(out io.Writer, lvl def.Level, f fmr.Formatter, registry *Logger) LoggerInterface {
	return &appender{
		registry:  registry,
		lvl:       lvl,
		formatter: f,
		out:       out,
	}
}

// Output writes the output for a logging event. The string s contains
// the text to print after the prefix specified by the formatter of the
// appender. A newline is appended if the last character of s is not
// already a newline.
func (a *appender) Output(s string) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.buf = a.buf[:0]
	a.buf = append(a.buf, a.formatter.GetPrefix(a.lvl)...)
	a.buf = append(a.buf, ": "...)
	a.buf = append(a.buf, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		a.buf = append(a.buf, '\n')
	}
	_, err := a.out.Write(a.buf)
	return err
}

// Printf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (a *appender) Printf(format string, v ...interface{}) {
	if a.registry.GetMaxLevel() < a.lvl {
		return
	}
	suffix := a.formatter.GetSuffix(a.lvl)
	format, v = a.formatter.Format(a.lvl, format, v...)
	_ = a.Output(fmt.Sprintf(format+suffix, v...))
}

// Print calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Print.
func (a *appender) Print(v ...interface{}) {
	if a.registry.GetMaxLevel() < a.lvl {
		return
	}
	_, v = a.formatter.Format(a.lvl, "", v...)
	v = append(v, a.formatter.GetSuffix(a.lvl))
	_ = a.Output(fmt.Sprint(v...))
}

// Println calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Println.
func (a *appender) Println(v ...interface{}) {
	if a.registry.GetMaxLevel() < a.lvl {
		return
	}
	_, v = a.formatter.Format(a.lvl, "", v...)
	v = append(v, a.formatter.GetSuffix(a.lvl))
	_ = a.Output(fmt.Sprintln(v...))
}

// Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
func (a *appender) Fatal(v ...interface{}) {
	if a.registry.GetMaxLevel() < a.lvl {
		return
	}
	_, v = a.formatter.Format(a.lvl, "", v...)
	v = append(v, a.formatter.GetSuffix(a.lvl))
	_ = a.Output(fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
func (a *appender) Fatalf(format string, v ...interface{}) {
	if a.registry.GetMaxLevel() < a.lvl {
		return
	}
	suffix := a.formatter.GetSuffix(a.lvl)
	format, v = a.formatter.Format(a.lvl, format, v...)
	_ = a.Output(fmt.Sprintf(format+suffix, v...))
	os.Exit(1)
}

// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
func (a *appender) Fatalln(v ...interface{}) {
	if a.registry.GetMaxLevel() < a.lvl {
		return
	}
	_, v = a.formatter.Format(a.lvl, "", v...)
	v = append(v, a.formatter.GetSuffix(a.lvl))
	_ = a.Output(fmt.Sprintln(v...))
	os.Exit(1)
}

// Panic is equivalent to l.Print() followed by a call to panic().
func (a *appender) Panic(v ...interface{}) {
	if a.registry.GetMaxLevel() < a.lvl {
		return
	}
	_, v = a.formatter.Format(a.lvl, "", v...)
	v = append(v, a.formatter.GetSuffix(a.lvl))
	s := fmt.Sprint(v...)
	_ = a.Output(s)
	panic(s)
}

// Panicf is equivalent to l.Printf() followed by a call to panic().
func (a *appender) Panicf(format string, v ...interface{}) {
	if a.registry.GetMaxLevel() < a.lvl {
		return
	}
	suffix := a.formatter.GetSuffix(a.lvl)
	format, v = a.formatter.Format(a.lvl, format, v...)
	s := fmt.Sprintf(format+suffix, v...)
	_ = a.Output(s)
	panic(s)
}

// Panicln is equivalent to l.Println() followed by a call to panic().
func (a *appender) Panicln(v ...interface{}) {
	if a.registry.GetMaxLevel() < a.lvl {
		return
	}
	_, v = a.formatter.Format(a.lvl, "", v...)
	v = append(v, a.formatter.GetSuffix(a.lvl))
	s := fmt.Sprintln(v...)
	_ = a.Output(s)
	panic(s)
}
