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