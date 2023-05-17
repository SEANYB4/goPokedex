package main

import (
	"fmt"
	"bufio"
	"os"
)


func main() {

	fmt.Println("type something")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()


	fmt.Println("echoing: ", text)
}