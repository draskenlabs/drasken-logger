package main

import "github.com/draskenlabs/drasken-logger/logger"

func main() {
	log := logger.New(logger.DEBUG, true)

	// Set to colorize just the log level (default)
	log.ColorTarget = "level"
	log.Info("Level color only")

	// Set to colorize just the message
	log.ColorTarget = "message"
	log.Warn("Message color only")

	// Set to colorize the entire line
	log.ColorTarget = "full"
	log.Error("Full line is colored")

	// Custom color config (optional)
	log.LevelColors[logger.INFO] = logger.ColorConfig{
		Prefix: "\033[95m", // Purple
		Suffix: "\033[0m",
	}
	log.ColorTarget = "message"
	log.Info("Now purple info messages")
}
