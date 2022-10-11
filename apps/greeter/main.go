package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ccccclong/nx-go-sandbox/libs/speaker/cat"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello! Whats your name?")
	name, _ := reader.ReadString('\n')
	fmt.Print(cat.Speak("Nice to meet you " + name))
}
