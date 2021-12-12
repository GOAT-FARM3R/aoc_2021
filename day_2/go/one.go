package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var horizontal, depth int = 0, 0

	reader := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for reader.Scan() {
		input := strings.Split(reader.Text(), " ")

		var value, _ = strconv.Atoi(input[1])

		command := input[0]

		switch {
			case command == "forward":
				horizontal += value
			case command == "down":
				depth += value
			case command == "up":
				depth -= value
		}
	}

	fmt.Printf("%d", horizontal * depth)
}