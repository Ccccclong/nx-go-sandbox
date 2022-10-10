package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello! Whats your name?")
	name, _ := reader.ReadString('\n')
	fmt.Println("Nice to meet you", name)
}
