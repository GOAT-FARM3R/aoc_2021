package main

import (
    "bufio"
    "fmt"
	"os"
	"bytes"
	"strconv"
)

func main() {

	inputs_zero := [] int {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	inputs_one := [] int {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	reader := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for reader.Scan() {
		var input = reader.Text()

		for pos, char := range input {
			if (string(char) == "0") {

				inputs_zero[pos] += 1
				continue
			}
			inputs_one[pos] += 1
		}

	}

	var output_gamma bytes.Buffer
	var output_epsilon bytes.Buffer

	for i:=0; i < len(inputs_one); i++ {
		if (inputs_one[i] > inputs_zero[i]) {
			output_gamma.WriteString("1")
			continue
		}
		output_gamma.WriteString("0")
	}

	for i:=0; i < len(inputs_one); i++ {
		if (inputs_zero[i] < inputs_one[i]) {
			output_epsilon.WriteString("0")
			continue
		}
		output_epsilon.WriteString("1")
	}

	i_gamma, _ := strconv.ParseInt(output_gamma.String(), 2, 64)
	i_epsilon, _ := strconv.ParseInt(output_epsilon.String(), 2, 64)

	fmt.Println((i_gamma*i_epsilon))
}