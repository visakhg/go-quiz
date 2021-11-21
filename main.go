package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	filename := "questions.csv"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error opening csv file '%s': %s", filename, err)
	}

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatalf("error reading csv file '%s': %s", file.Name(), err)
	}

	//results := make(map[string]bool, len(records))
	score := 0
	answer := ""
	fmt.Println("Starting quiz...")
	time.Sleep(1 * time.Second)
	for _, question := range records {
		fmt.Println(question[0])
		fmt.Scanf("%s", &answer)
		if answer == question[1] {
			score++
		}
	}

	fmt.Printf("\nYour results are being calculated... \n")
	time.Sleep(1 * time.Second)
	fmt.Printf("Score: %d/%d\n", score, len(records))
}
