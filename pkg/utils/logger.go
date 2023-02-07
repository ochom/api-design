package utils

import (
	"fmt"
	"log"
	"os"
)

// CustomLogger is a custom logger
type CustomLogger struct {
	k *log.Logger
}

// NewLogger creates a new logger instance
func NewLogger() *CustomLogger {
	customLogger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	return &CustomLogger{k: customLogger}
}

func (c *CustomLogger) print(txt string) {
	_ = c.k.Output(3, txt)
}

// Info prints informational messages
func (c *CustomLogger) Info(v ...any) {
	c.print("INFO: " + fmt.Sprint(v...))
}

// Infof prints informational messages with formatting
func (c *CustomLogger) Infof(format string, v ...any) {
	c.print("INFO: " + fmt.Sprintf(format, v...))
}

// Warn  prints warning messages
func (c *CustomLogger) Warn(v ...any) {
	c.print("WARN: " + fmt.Sprint(v...))
}

// Warnf prints warning messages with formatting
func (c *CustomLogger) Warnf(format string, v ...any) {
	c.print("WARN: " + fmt.Sprintf(format, v...))
}

// Error prints error messages
func (c *CustomLogger) Error(v ...any) {
	c.print("ERROR: " + fmt.Sprint(v...))
}

// Errorf prints error messages with formatting
func (c *CustomLogger) Errorf(format string, v ...any) {
	c.print("ERROR: " + fmt.Sprintf(format, v...))
}

// Fatal prints fatal messages and exits
func (c *CustomLogger) Fatal(v ...any) {
	c.print("FATAL: " + fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf prints fatal messages with formatting and exits
func (c *CustomLogger) Fatalf(format string, v ...any) {
	c.print("FATAL: " + fmt.Sprintf(format, v...))
	os.Exit(1)
}
