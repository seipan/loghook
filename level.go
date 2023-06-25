package loghook

import (
	"errors"
	"fmt"
	"strings"
)

// Level is a logging priority. Higher levels are more important.
type Level uint32

const (
	// Debug level logs are used for debugging
	DebugLevel Level = iota
	// Info level logs are logs at a lower level and are used to preserve information.
	InfoLevel
	// Warn level is a higher level of logging than Info level
	// and is usually used to output more important logs than the info log.
	WarnLevel
	// Error level is a higher level of logging and is usually used to output a log of errors
	ErrorLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel

	// In this logger, the lowest log level is Debug level and the highest level is Fatal Level
	_minLevel = DebugLevel
	_maxLevel = FatalLevel
)

// Converts log level to string, returns unknown if log level is not expected.
func (l Level) String() string {
	str := l.MarshalText()
	return str
}

func (l Level) MarshalText() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case PanicLevel:
		return "panic"
	case FatalLevel:
		return "fatal"
	default:
		return "unknown"
	}
}

func (l Level) UppercaseString() string {
	str := l.String()
	return strings.ToUpper(str)
}

// Converts the string representing level to level. Returns an error if the level is not expected
func (l *Level) UnmarshalText(text string) error {
	if l == nil {
		return errors.New("level nil")
	}

	lvl, err := ParseLevel(text)
	if err != nil {
		return err
	}

	*l = lvl
	return nil
}

func ParseLevel(lvl string) (Level, error) {
	switch strings.ToLower(lvl) {
	case "debug":
		return DebugLevel, nil
	case "info":
		return InfoLevel, nil
	case "warn", "warning":
		return WarnLevel, nil
	case "error":
		return ErrorLevel, nil
	case "panic":
		return PanicLevel, nil
	case "fatal":
		return FatalLevel, nil
	}

	var l Level
	return l, fmt.Errorf("not a valid logdis Level: %q", lvl)
}

// Converts from string to level and sets
func (l *Level) Set(s string) error {
	return l.UnmarshalText(s)
}

// Get the level
func (l *Level) Get() interface{} {
	return *l
}
