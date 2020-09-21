package main

import (
	"fmt"
	"os"
	"os/exec"
)

func exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func delFileIfExists(file string) {
	if exists(file) {
		errIt(os.Remove(file), "")
	}
}

func openFileInEditor(filename string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	executable, err := exec.LookPath(editor)
	errIt(err, "")
	cmd := exec.Command(executable, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func errIt(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}

func createFile(file string, content string) {
	delFileIfExists(file)
	f, err := os.Create(file)
	errIt(err, "")
	defer f.Close()
	_, err = f.WriteString(content)
	errIt(err, "")
}
