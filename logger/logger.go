package logger

import (
	"fmt"
	"os"
	"time"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
)

var defaultLevelNames = []string{"DEBUG", "INFO", "WARN", "ERROR"}

type ColorConfig struct {
	Prefix string // Start of color
	Suffix string // End of color
}

type Logger struct {
	MinLevel    int
	UseColor    bool
	LevelNames  []string
	LevelColors []ColorConfig // One per level
	ColorTarget string        // "level", "message", "full"
}

// New creates a new logger with optional color and default color config.
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
		ColorTarget: "level", // or "message" or "full"
	}
}

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

	// No full-line color or color disabled
	fmt.Fprintf(os.Stdout, "[%s] [%s] %s\n", timestamp, levelName, msg)
}

// Shortcut methods
func (l *Logger) Debug(format string, args ...interface{}) { l.log(DEBUG, format, args...) }
func (l *Logger) Info(format string, args ...interface{})  { l.log(INFO, format, args...) }
func (l *Logger) Warn(format string, args ...interface{})  { l.log(WARN, format, args...) }
func (l *Logger) Error(format string, args ...interface{}) { l.log(ERROR, format, args...) }
