package day_2

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func getFileContentIntFormat(fileName string) (depth int, horizontal int) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	lines := strings.Split(text, "\n")

	depth = 0
	horizontal = 0
	aim := 0

	for _, line := range lines {
		directionAndRange := strings.Split(line, " ")
		parsedRange, _ := strconv.ParseInt(directionAndRange[1], 10, 64)
		direction := directionAndRange[0]

		switch direction {
		case "forward":
			horizontal = horizontal + int(parsedRange)
			depth = depth + (aim * int(parsedRange))
		case "down":
			aim = aim + int(parsedRange)
		case "up":
			aim = aim - int(parsedRange)
		}

	}

	return
}
