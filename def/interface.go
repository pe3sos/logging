package def

// Level type
type Level int

var Levels = struct {
	FATAL   Level
	ERROR   Level
	WARNING Level
	INFO    Level
	DEBUG   Level
}{
	FATAL:   0,
	ERROR:   1,
	WARNING: 2,
	INFO:    3,
	DEBUG:   4,
}

var names = map[Level]string{
	Levels.FATAL:   "FATAL",
	Levels.ERROR:   "ERROR",
	Levels.WARNING: "WARNING",
	Levels.INFO:    "INFO",
	Levels.DEBUG:   "DEBUG",
}

func (l Level) Name() string {
	return names[l]
}

// Formatter interface
type Formatter interface {
	GetPrefix(lvl Level) string
	Format(lvl Level, format string, values ...interface{}) (string, []interface{})
	GetSuffix(lvl Level) string
}

// LoggerInterface will accept stdlib logger and a custom logger.
// There's no standard interface, this is the closest we get, unfortunately.
type LoggerInterface interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
}
