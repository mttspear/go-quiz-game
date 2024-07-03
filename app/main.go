package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var correct int = 0

func main() {
	// the below will out put when we run with the -h flag  -csv string a csv file in the format of 'question, answer' (default "problems.csv")
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")

	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q) //print out question plus one so quiz starts at one instead of 0
		answerCh := make(chan string)
		//go routine
		go func() {
			var answer string
			fmt.Scanf("%s \n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d. \n", correct, len(problems))
			return
		case answer := <-answerCh:
			ScoreAnswer(answer, p.a)
		}
	}

	fmt.Printf("You scored %d out of %d. \n", correct, len(problems))
}

//Score the answer
func ScoreAnswer(userAnswer string, questionAnswer string) bool {
	if userAnswer == questionAnswer {
		correct++
		fmt.Println("Correct!")
		return true
	} else {
		fmt.Println("Erroneous!")
		return false
	}
}

//Parse the csv in to a struct
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
