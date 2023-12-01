package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Hello world, the web server
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	re := regexp.MustCompile(`\d`)

	for scanner.Scan() {
		digitsInLine := re.FindAllString(scanner.Text(), -1)
		firstDigit := digitsInLine[0]
		lastDigit := digitsInLine[len(digitsInLine)-1]
		combinedValueStr := firstDigit + lastDigit
		combinedValue, err := strconv.Atoi(combinedValueStr)
		if err != nil {
			// ... handle error
			panic(err)
		}
		sum = sum + combinedValue
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
