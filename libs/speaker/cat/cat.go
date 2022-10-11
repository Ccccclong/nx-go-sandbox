package cat

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func Speak(message string, times int) string {
	dirname := getDirname()
	fileContent, err := os.ReadFile(dirname + "/cat.txt")
	if err != nil {
		panic(err)
	}
	image := fmt.Sprintf(string(fileContent), message) + "\n"
	images := strings.Repeat(image, times)
	return images
}

func getDirname() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("unable to get the current filename"))
	}
	dirname := filepath.Dir(filename)
	return dirname
}
