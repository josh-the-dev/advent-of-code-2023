package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func cubeAmountCheck(cubeString string) bool {
	redLimit := 12
	greenLimit := 13
	blueLimit := 14

	cubeAmount := strings.Fields(cubeString)[0]
	cubeAmountInt, err := strconv.Atoi(cubeAmount)
	if err == nil {
		if strings.Contains(cubeString, "red") {
			return cubeAmountInt > redLimit

		}
		if strings.Contains(cubeString, "blue") {
			return cubeAmountInt > blueLimit
		}
		if strings.Contains(cubeString, "green") {

			return cubeAmountInt > greenLimit

		}
	}
	return false
}

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numberRegex := regexp.MustCompile("[0-9]+")
	sum := 0
	for scanner.Scan() {
		rowSplit := strings.Split(scanner.Text(), ":")
		rowId := numberRegex.FindString(rowSplit[0])
		rowSubsets := strings.Split(rowSplit[1], ";")

		subsetFailure := false
		for i, subset := range rowSubsets {
			_ = i
			colourSplit := strings.Split(subset, ",")
			for j, colourAmount := range colourSplit {
				_ = j
				if !subsetFailure {
					subsetFailure = cubeAmountCheck(colourAmount)
				}
			}
		}

		if !subsetFailure {
			rowIdInt, err := strconv.Atoi(rowId)
			fmt.Println(rowIdInt)
			if err == nil {
				sum = sum + rowIdInt
			}
		}

	}

	fmt.Println(sum)

}
