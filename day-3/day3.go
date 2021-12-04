package day_3

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func getBitMap(fileName string) (bits map[int][]int, sortedKeys []int) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	bits = map[int][]int{}
	text := string(content)
	lines := strings.Split(text, "\n")

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

	sortedKeys = []int{}

	for key, _ := range bits {
		sortedKeys = append(sortedKeys, key)
	}

	sort.Ints(sortedKeys)

	return
}

func calculateOxigenRate() (oxigenRate string) {
	bits, bitsKeys := getBitMap("day3.txt")
	rowsToIgnore := []int{}

	oxigenRate = ""

	for _, key := range bitsKeys {
		column := bits[key]

		zeros := []int{}
		ones := []int{}

		for iterator, bit := range column {
			shouldIgnore := false

			for _, number := range rowsToIgnore {
				if iterator == number {
					shouldIgnore = true
					break
				}
			}

			if shouldIgnore {
				continue
			}

			if bit == 0 {
				zeros = append(zeros, iterator)
			} else {
				ones = append(ones, iterator)
			}
		}

		if len(ones) > len(zeros) || len(ones) == len(zeros) {
			for _, toIgnore := range zeros {
				rowsToIgnore = append(rowsToIgnore, toIgnore)
			}
			oxigenRate += "1"
		} else {
			for _, toIgnore := range ones {
				rowsToIgnore = append(rowsToIgnore, toIgnore)
			}
			oxigenRate += "0"
		}
	}

	return
}

func calculateCo2Scrubber() (co2Scrubber string) {
	bits, bitsKeys := getBitMap("day3.txt")
	rowsToIgnore := []int{}

	co2Scrubber = ""

	for _, key := range bitsKeys {
		column := bits[key]

		zeros := []int{}
		ones := []int{}

		for iterator, bit := range column {
			shouldIgnore := false

			for _, number := range rowsToIgnore {
				if iterator == number {
					shouldIgnore = true
					break
				}
			}

			if shouldIgnore {
				continue
			}

			if len(rowsToIgnore)+1 == len(column) {
				result := ""

				for _, keyOfBit := range bitsKeys {
					result += strconv.Itoa(bits[keyOfBit][iterator])
				}

				return result
			}

			if bit == 0 {
				zeros = append(zeros, iterator)
			} else {
				ones = append(ones, iterator)
			}
		}

		if len(ones) > len(zeros) || len(ones) == len(zeros) {
			for _, toIgnore := range ones {
				rowsToIgnore = append(rowsToIgnore, toIgnore)
			}
			co2Scrubber += "0"
		} else {
			for _, toIgnore := range zeros {
				rowsToIgnore = append(rowsToIgnore, toIgnore)
			}
			co2Scrubber += "1"
		}
	}

	return
}

func gatherRates() (gammaRate string, epsilonRate string) {
	bits, bitsKeys := getBitMap("day3.txt")

	gammaRate = ""
	epsilonRate = ""

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
	return gammaRate, epsilonRate
}

func solutions() {
	oxigenRate := calculateOxigenRate()
	co2Scrubber := calculateCo2Scrubber()

	fmt.Println(co2Scrubber)
	convertedOxigenRate, _ := strconv.ParseInt(oxigenRate, 2, 64)
	fmt.Println(convertedOxigenRate)
	convertedCo2Scrubber, _ := strconv.ParseInt(co2Scrubber, 2, 64)
	fmt.Println(convertedCo2Scrubber)
	fmt.Println(convertedOxigenRate * convertedCo2Scrubber)

	/*gammaRate, epsilonRate := getBitMap("day3.txt")

	convertedGamma, _ := strconv.ParseInt(gammaRate, 2, 64)
	convertedEpsilon, _ := strconv.ParseInt(epsilonRate, 2, 64)

	fmt.Println(gammaRate, epsilonRate)
	fmt.Println(convertedGamma, convertedEpsilon)
	fmt.Println(convertedGamma * convertedEpsilon)*/
}
