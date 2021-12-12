package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
)

func main() {

	var increased_counter int = 0
	var window []int

	reader := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for reader.Scan() {
		var current, _ = strconv.Atoi(reader.Text())
		window = append(window, current)

		if (len(window) != 4) {
			continue
		}

		var a,b int = 0,0

		for _,input := range window[0:3]{ 
			a+=input
		}

		for _,input := range window[1:4]{ 
			b+=input
		}

		if (b > a) {
			increased_counter++
		}
		window = append(window[:0], window[1:]...)
	}

	fmt.Printf("%d", increased_counter)
}