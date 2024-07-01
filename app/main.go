package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// the below will out put when we run with the -h flag  -csv string a csv file in the format of 'question, answer' (default "problems.csv")
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		//catch any err opening the file
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)

	correct := 0

	for i, p := range problems {
		//Below can be broken into own function
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q) // plus one so quiz starts at one instead of 0
		var answer string
		fmt.Scanf("%s \n", &answer)
		if answer == p.a {
			correct++
			fmt.Println("Correct!")
		}
	}

	fmt.Printf("You scored %d out of %d. \n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: strings.TrimSpace(line[0]), // first column is question
			a: strings.TrimSpace(line[1]), // second column is the answer
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1) // exit the application with status of 1
}
