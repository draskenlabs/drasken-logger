# drasken-logger

A lightweight, customizable logger for Go with **colorized terminal output** and **log level control**.

## ✨ Features

- Log levels: `DEBUG`, `INFO`, `WARN`, `ERROR`
- Color output in terminal (optional)
- Customizable:
  - Where to apply color: level, message, or full line
  - Per-level color codes
  - Min log level filtering

## 📦 Installation

```bash
go get github.com/draskenlabs/drasken-logger
```

## 🔧 Usage

```go
package main

import "github.com/draskenlabs/drasken-logger"

func main() {
	log := logger.New(logger.DEBUG, true)

	log.Debug("Debugging %v", "something")
	log.Info("Server started on port %d", 8080)
	log.Warn("Memory usage is high")
	log.Error("Unable to connect to DB")
}
```

## 🎨 Color Modes

```go
log.ColorTarget = "level"   // (default) color only the log level
log.ColorTarget = "message" // color only the message
log.ColorTarget = "full"    // color the entire log line
```

## 🎯 Customize Colors

```go
log.LevelColors[logger.INFO] = logger.ColorConfig{
    Prefix: "\033[95m", // Purple
    Suffix: "\033[0m",
}
```

## 📘 Filter by Log Level

```go
log := logger.New(logger.WARN, true)
// Only WARN and ERROR logs will be printed
```

## 🖥️ Example Output

```go
[2025-04-06 14:35:02] [DEBUG] Starting debug mode
[2025-04-06 14:35:03] [INFO] Server running at :8080
[2025-04-06 14:35:04] [WARN] Memory usage high
[2025-04-06 14:35:05] [ERROR] Failed to connect to DB
```

## 📜 License

MIT © [Drasken Labs](https://github.com/draskenlabs)


---

Want me to generate a basic project structure (`go.mod`, `main.go`, `logger/logger.go`, etc.) and push-ready repo layout next?
