package middlewareHelpers

import (
    "log"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "os"
)

func SetupLogger(e *echo.Echo, logPath string) {
	if len(logPath) > 0 {
		logFile, err := os.Create(logPath)
		if err == nil {
			defer logFile.Close()
			e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
				Output: logFile,
			}))
			log.Printf("INFO: Logging API calls to %s\n", logPath)
		} else {
			log.Println("WARN: Unable to create log file!")
		}
	} else {
		log.Println("WARN: Logging disabled")
	}
}
