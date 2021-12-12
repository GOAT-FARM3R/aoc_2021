package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var horizontal, depth, aim int = 0, 0, 0

	reader := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for reader.Scan() {
		input := strings.Split(reader.Text(), " ")

		var value, _ = strconv.Atoi(input[1])

		command := input[0]

		switch {
			case command == "forward":
				horizontal += value
				depth += (aim * value)
			case command == "down":
				aim += value
			case command == "up":
				aim -= value
		}
	}

	fmt.Printf("%d", horizontal * depth)
}