package cat

import (
	"fmt"
	"os"
)

func Speak(message string) string {
	content, _ := os.ReadFile("cat.txt")
	return fmt.Sprintf(string(content), message)
}
