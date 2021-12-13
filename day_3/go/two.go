package main

import (
    "bufio"
    "fmt"
	"os"
	//"bytes"
	"strconv"
)

func main() {

	reader := bufio.NewScanner(bufio.NewReader(os.Stdin))

	var oxygen, cotwo [] string

	for reader.Scan() {
		var input = reader.Text()

		oxygen = append(oxygen, input)
	}

	cotwo = oxygen

	pos := 0

	for true {

		if ((len(oxygen) == 1 && len(cotwo) == 1)) {
			break
		}

		var ones, zeros [] string

		if (len(oxygen) != 1) {

			for _, input := range oxygen {
				if (string(input[pos]) ==  "1") {
					ones = append(ones, input)
					continue
				}
				zeros = append(zeros, input)
			}

			if (len(ones) > len(zeros)) {
				oxygen = ones
			} else if (len(ones) == len(zeros)) {
				oxygen = ones
			} else {
				oxygen = zeros
			}
		}

		if (len(cotwo) != 1) {

			ones = nil 
			zeros = nil

			for _, input := range cotwo {
				if (string(input[pos]) ==  "0") {
					zeros = append(zeros, input)
					continue
				}
				ones = append(ones, input)
			}
			
			if (len(zeros) < len(ones)) {
				cotwo = zeros
			} else if (len(ones) == len(zeros)) {
				cotwo = zeros
			} else {
				cotwo = ones
			}
		}
		pos += 1

	}

	i_oxygen, _ := strconv.ParseInt(string(oxygen[0]), 2, 64)
	i_cotwo, _ := strconv.ParseInt(string(cotwo[0]), 2, 64)

	fmt.Println((i_oxygen*i_cotwo))
}