// Package clio provides CLI I/O utilities.
package clio

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

// ReadPw reads a password from file or stdin prompt.
func ReadPw(file *os.File, out io.Writer) (string, error) {
	if file != nil && fileHasData(file) {
		reader := bufio.NewReader(file)
		text, err := reader.ReadString('\n')
		return strings.Trim(text, "\n\r"), err
	}
	fmt.Fprint(out, "Password:")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	return string(bytePassword), err
}

func fileHasData(file *os.File) bool {
	fi, err := file.Stat()
	return err == nil && fi.Size() > 0
}
