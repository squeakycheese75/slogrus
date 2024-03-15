# slogrus

`slogrus` is a this simple wrapper that allows me to use `slog` but in a more `logrus` way. 

## Installation

```sh
go get -u github.com/squeakycheese75/slogrus
```

## Usage 
```go
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

	// Log a debug with fields
	logger.WithFields(slogrus.Fields{
		"http":          "req",
		"method":        "POST",
		"handler":       "UpdateDataColumn",
		"correlationID": "order123",
		"params": map[string]string{
			"datacolumnID": "abcdef",
			"datasetId":    "123456",
		},
	}).Debug("UpdateDataColumn request received")
}
```

## Contributing
Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are greatly appreciated.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".

Don't forget to give the project a star! Thank you again!

Fork the Project
Create your Feature Branch (git checkout -b feature/AmazingFeature)
Commit your Changes (git commit -m 'Add some AmazingFeature')
Push to the Branch (git push origin feature/AmazingFeature)
Open a Pull Request
License
Distributed under the MIT License. See LICENSE for more information.

## Acknowledgments
- logrus for the inspiration.
- slog for the performance-focused logging.