package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	gammaRate, epsilonRate := getFileContentIntFormat("day3.txt")

	convertedGamma, _ := strconv.ParseInt(gammaRate, 2, 64)
	convertedEpsilon, _ := strconv.ParseInt(epsilonRate, 2, 64)

	fmt.Println(gammaRate, epsilonRate)
	fmt.Println(convertedGamma, convertedEpsilon)
	fmt.Println(convertedGamma * convertedEpsilon)
}

func getFileContentIntFormat(fileName string) (gammaRate string, epsilonRate string) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	lines := strings.Split(text, "\n")

	gammaRate = ""
	epsilonRate = ""
	bits := map[int][]int{}

	for _, line := range lines {
		for iterator, character := range line {
			parsedInt, _ := strconv.ParseInt(string(character), 10, 64)

			if bits[iterator] == nil {
				bits[iterator] = []int{int(parsedInt)}
			} else {
				bits[iterator] = append(bits[iterator], int(parsedInt))
			}
		}
	}

	bitsKeys := []int{}

	for key, _ := range bits {
		bitsKeys = append(bitsKeys, key)
	}

	sort.Ints(bitsKeys)

	for _, key := range bitsKeys {
		column := bits[key]

		zeros := 0
		ones := 0

		for _, bit := range column {
			if bit == 0 {
				zeros++
			} else {
				ones++
			}
		}

		if ones > zeros {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	return
}
