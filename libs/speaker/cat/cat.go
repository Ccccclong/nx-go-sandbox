package cat

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func Speak(message string) string {
	dirname := getDirname()
	content, err := os.ReadFile(dirname + "/cat.txt")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf(string(content), message)
}

func getDirname() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("unable to get the current filename"))
	}
	dirname := filepath.Dir(filename)
	return dirname
}
