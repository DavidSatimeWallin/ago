package util

import (
	"fmt"
	"os"
	"os/exec"
)

// FileExists returns true of a file exists and is not a dir
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// FolderExists returns true if the folder exists and is not a file
func FolderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func DelFileIfExists(file string) {
	if FileExists(file) {
		err := os.Remove(file)
		ErrIt(err, "")
	}
}

// DefaultEditor will be vim cause that's what real adults use
const DefaultEditor = "vim"

// OpenFileInEditor opens filename in a text editor.
func OpenFileInEditor(filename string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = DefaultEditor
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
	defer f.Close()
	_, err = f.WriteString(content)
	ErrIt(err, "")
}
