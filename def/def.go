package def

// Level type
type Level int

const (
	FATAL Level = iota
	ERROR
	WARNING
	INFO
	DEBUG
)
