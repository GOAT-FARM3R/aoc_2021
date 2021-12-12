package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
)

func main() {

	var increased_counter, previous int = 0, -1

	reader := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for reader.Scan() {
		var current, _ = strconv.Atoi(reader.Text())

		if (previous == -1) {
			previous = current
			continue
		}

		if (current > previous) {
			increased_counter++
		}
		previous = current
	}

	fmt.Printf("%d", increased_counter)
}