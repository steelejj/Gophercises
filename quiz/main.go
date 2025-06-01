package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Question struct {
	question string
	answer   string
}

func loadQuestions(file *os.File) ([]Question, error) {
	reader := csv.NewReader(file)

	var questions []Question

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		questions = append(questions, Question{row[0], row[1]})
	}

	return questions, nil
}

func answerQuestion() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter your answer")
	if scanner.Scan() {
		return strings.Trim(scanner.Text(), " ")
	} else {
		return ""
	}
}

func main() {

	fmt.Println(os.Getwd())

	f, err := os.Open("quiz/p2.csv") // replace with your file path
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer f.Close()

	questions, err := loadQuestions(f)

	if err != nil {
		log.Fatalf("failed to load questions: %v", err)
	}

	var correct int

	for _, q := range questions {
		fmt.Println("Question: ", q.question)
		answer := answerQuestion()

		if answer == q.answer {
			correct += 1
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect!")
		}
	}

	fmt.Printf("You got %d correct out of %d\n", correct, len(questions))

}
