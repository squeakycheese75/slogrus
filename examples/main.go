package main

import (
	"github.com/squeakycheese75/slogrus"
)

func main() {
	logger := slogrus.New()

	// Log an info message
	logger.Info("This is an info message with slogrus")

	// Log an error message with fields
	logger.WithFields(slogrus.Fields{
		"error": "some error message",
	}).Error("An error occurred")

	// Log some json with fields
	logger.WithFields(slogrus.Fields{
		"http":          "req",
		"method":        "PATCH",
		"handler":       "UpdateDataColumn",
		"correlationID": "order123",
		"params": map[string]string{
			"datacolumnID": "abcdef",
			"datasetId":    "123456",
		},
	}).Debug("UpdateDataColumn request received")
}
