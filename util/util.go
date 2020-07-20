package util

import (
	"fmt"
	"os"
	"os/exec"
)

// Exists returns true if a file or folder exists
func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func DelFileIfExists(file string) {
	if Exists(file) {
		err := os.Remove(file)
		ErrIt(err, "")
	}
}

// OpenFileInEditor opens filename in a text editor.
func OpenFileInEditor(filename string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	executable, err := exec.LookPath(editor)
	ErrIt(err, "")
	cmd := exec.Command(executable, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// ErrIt is an ugly way of handling errors
func ErrIt(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}

func GenerateFile(file string, content string) {
	DelFileIfExists(file)
	f, err := os.Create(file)
	ErrIt(err, "")
	defer f.Close()
	_, err = f.WriteString(content)
	ErrIt(err, "")
}
