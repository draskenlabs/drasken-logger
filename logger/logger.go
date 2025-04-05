package logger

import (
	"fmt"
	"os"
	"time"
)

// Log level constants for filtering
const (
	DEBUG = iota
	INFO
	WARN
	ERROR
)

// Default log level names (used if not overridden)
var defaultLevelNames = []string{"DEBUG", "INFO", "WARN", "ERROR"}

// ColorConfig represents ANSI color codes to wrap around log parts.
type ColorConfig struct {
	Prefix string // Start of color (e.g., \033[36m)
	Suffix string // Reset color (e.g., \033[0m)
}

// Logger provides configurable logging with levels and optional color.
type Logger struct {
	MinLevel    int           // Minimum level to log (e.g., INFO and above)
	UseColor    bool          // Whether to use color output
	LevelNames  []string      // Names for each log level
	LevelColors []ColorConfig // ANSI color configs per level
	ColorTarget string        // Which part to color: "level", "message", or "full"
}

// New creates and returns a new Logger instance with default colors.
func New(minLevel int, useColor bool) *Logger {
	return &Logger{
		MinLevel:   minLevel,
		UseColor:   useColor,
		LevelNames: defaultLevelNames,
		LevelColors: []ColorConfig{
			{Prefix: "\033[36m", Suffix: "\033[0m"}, // DEBUG - Cyan
			{Prefix: "\033[32m", Suffix: "\033[0m"}, // INFO - Green
			{Prefix: "\033[33m", Suffix: "\033[0m"}, // WARN - Yellow
			{Prefix: "\033[31m", Suffix: "\033[0m"}, // ERROR - Red
		},
		ColorTarget: "level", // Options: "level", "message", or "full"
	}
}

// log prints a formatted log message based on level, color, and configuration.
func (l *Logger) log(level int, format string, args ...interface{}) {
	if level < l.MinLevel {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf(format, args...)
	levelName := l.LevelNames[level]
	color := l.LevelColors[level]

	switch l.ColorTarget {
	case "level":
		if l.UseColor {
			levelName = color.Prefix + levelName + color.Suffix
		}
	case "message":
		if l.UseColor {
			msg = color.Prefix + msg + color.Suffix
		}
	case "full":
		if l.UseColor {
			fmt.Fprintf(os.Stdout, "%s[%s] [%s] %s%s\n", color.Prefix, timestamp, levelName, msg, color.Suffix)
			return
		}
	}

	// Default output (no full-line color or color disabled)
	fmt.Fprintf(os.Stdout, "[%s] [%s] %s\n", timestamp, levelName, msg)
}

// Debug logs a message at DEBUG level.
func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(DEBUG, format, args...)
}

// Info logs a message at INFO level.
func (l *Logger) Info(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}

// Warn logs a message at WARN level.
func (l *Logger) Warn(format string, args ...interface{}) {
	l.log(WARN, format, args...)
}

// Error logs a message at ERROR level.
func (l *Logger) Error(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
}
