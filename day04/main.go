package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	sum, err := readInputFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum:", sum)
}

func readInputFile(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	runningTotal := 0
	// digitRegex := regexp.MustCompile(`\d+`)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, "|")
		winningNumbersWithCardTitle := lineSplit[0]
		winningNumbers := strings.Split(winningNumbersWithCardTitle, ":")
		winningNumbers = strings.Fields(winningNumbers[1])

		if err != nil {
			fmt.Println(err)
		}
		drawnNumbers := strings.Fields(lineSplit[1])
		matches := 0
		for _, drawnNumber := range drawnNumbers {
			if contains(winningNumbers, drawnNumber) {
				if matches == 0 {
					matches += 1
				} else {
					matches *= 2
				}
			}
		}
		runningTotal += matches
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	return runningTotal, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
