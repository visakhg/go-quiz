package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type problem struct{
	question string
	answer string
}

func main() {
	filename := flag.String("csv", "questions.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "time limit for the quiz in seconds")
	flag.Parse()

	records, err := readCsv(*filename)
	if err != nil {
		log.Fatalf("error reading csv file '%s': %s", *filename, err)
	}
	
	problems := parseRecords(records)
	stopwatch := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	score := 0
	answer := ""

	fmt.Println("Starting quiz...")
	time.Sleep(1 * time.Second)
	problemloop:
		for i, problem := range problems {
			fmt.Printf("\nProblem #%d\n", i+1)
			fmt.Println(problem.question)

			answerChan := make(chan string)
			go func(answer string) {
				fmt.Scanf("%s\n", &answer)
				answerChan <- answer
			}(answer)

			select {
			case <-stopwatch.C:
				fmt.Printf("\nYour time is up!")
				break problemloop
			case answer := <-answerChan:
				if answer == problem.answer {
					score++
				}
			}
		}

	fmt.Printf("\nYour results are being calculated... \n")
	time.Sleep(1 * time.Second)
	fmt.Printf("Score: %d/%d\n", score, len(records))
}

func parseRecords(records [][]string) []problem {
	problems := make([]problem, len(records))
	for i, record := range records {
		problems[i] = problem{question: record[0], answer: strings.TrimSpace(record[1])}
	}
	return problems
}

func readCsv(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}