package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
	var line string
	var array []string
	var arraySize int
	_ = arraySize

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	arraySize, err := strconv.Atoi(scanner.Text());
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot parse input:", err)
		os.Exit(1)
	}

	scanner.Scan()
	line = scanner.Text()

	array = strings.Split(line, " ")
	for i := len(array)-1; i >= 0; i-- {
		fmt.Printf("%s ", array[i])
	}
}