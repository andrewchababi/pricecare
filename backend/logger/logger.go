package logger

import (
	"io"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

// TODO make custom logger and logger middleware to keep history

type FilteredWriter struct {
	Writer io.Writer
}

func (fw *FilteredWriter) Write(p []byte) (n int, err error) {
	if strings.Contains(string(p), "http server started") {
		log.Printf("Server started")
		return 0, nil
	}
	return fw.Writer.Write(p)
}

func RegisterLogger(e *echo.Echo) {
	e.HideBanner = true
	e.Logger.SetOutput(&FilteredWriter{Writer: e.Logger.Output()})
}
