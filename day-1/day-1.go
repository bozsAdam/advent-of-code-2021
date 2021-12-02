package day_1

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func secondSolution() {
	numbers := getFileContentIntFormat("day1-1.txt")

	sums := []int64{}
	currentSum := 0
	sumsAdded := 0

	for i := 0; i < len(numbers); i++ {
		number := numbers[i]
		currentSum = currentSum + int(number)
		sumsAdded++

		if sumsAdded == 3 {
			sums = append(sums, int64(currentSum))
			sumsAdded = 0
			currentSum = 0
			i = i - 2
		}
	}

	firstSolution(sums)
}

func firstSolution(fileContent []int64) {
	previous := fileContent[0]
	counter := 0

	for i := 0; i < len(fileContent); i++ {
		current := fileContent[i]
		if current > previous {
			fmt.Printf("Current: %s increased \n", current)
			counter += 1
		} else {
			fmt.Printf("Current: %s decreased \n", current)
		}

		previous = current
	}

	fmt.Printf("All depth increases %s", counter)
}

func executeFirstSolution() {
	firstSolution(getFileContentIntFormat("day1-1.txt"))
}

func secondSolutionWithDataMarked() {
	sums := getFileContentHashMap("day1-2.txt")
	depths := []int64{}

	for _, value := range sums {
		depths = append(depths, value)
	}

	firstSolution(depths)
}

func getFileContentIntFormat(fileName string) []int64 {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	lines := strings.Split(text, "\n")
	depths := []int64{}

	for _, line := range lines {
		i, _ := strconv.ParseInt(line, 10, 64)
		depths = append(depths, i)
	}

	return depths
}

func getFileContentHashMap(fileName string) map[string]int64 {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	lines := strings.Split(text, "\n")

	rows := map[string][]int64{}

	for _, line := range lines {
		splittedLine := strings.Split(line, " ")
		currNumber, _ := strconv.ParseInt(splittedLine[0], 10, 64)

		for _, character := range splittedLine[1:] {
			if strings.TrimSpace(character) == "" {
				continue
			}

			character = strings.TrimSpace(character)
			numsInRow := rows[character]

			if numsInRow != nil {
				rows[character] = append(numsInRow, currNumber)
			} else {
				rows[character] = []int64{currNumber}
			}
		}
	}

	summarizedRows := map[string]int64{}

	for key, value := range rows {
		summarizedRows[key] = sum(value)
	}

	return summarizedRows
}

func sum(array []int64) int64 {
	result := int64(0)
	for _, v := range array {
		result += v
	}
	return result
}
