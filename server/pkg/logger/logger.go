package logger

import (
	"log/slog"
	"os"
)

// @Summary 		Logger Init
// @Description 	Used for the logging message and modules
// @Tags 			Log
// @Accept 			*/*
// @Produce 		plain
func LoggerInit() {
	jsonFormatter := slog.NewJSONHandler(os.Stdout, nil)
	slog.SetDefault(slog.New(jsonFormatter))
}
