package recoder

import (
	"log"
	"os"
)

// Write method is write error in error.log
func Write(errorMessage error) {
	if errorMessage != nil {
		file, fileError := os.OpenFile("log/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if fileError != nil {
			log.Fatalln(fileError)
		}
		defer file.Close()
		_, fileWriteError := file.WriteString(errorMessage.Error())
		if fileWriteError != nil {
			log.Fatalln(fileWriteError)
		}
	}
}
