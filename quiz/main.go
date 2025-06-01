package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Question struct {
	question string
	answer   int
}

type Quiz struct {
	questions []Question
}

func parseQuestion(record []string) (Question, error) {
	question := record[0]
	answer, err := strconv.Atoi(record[1])
	if err != nil {
		return Question{}, err
	}
	return Question{question: question, answer: answer}, nil
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
		question, err := parseQuestion(row)
		if err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}

	return questions, nil
}

func answerQuestion() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter your answer")
	if scanner.Scan() {
		return strconv.Atoi(scanner.Text())
	} else {
		return 0, fmt.Errorf("Failed to read input")
	}
}

func main() {

	fmt.Println(os.Getwd())

	f, err := os.Open("quiz/problems.csv") // replace with your file path
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer f.Close()

	questions, err := loadQuestions(f)

	for _, q := range questions {
		fmt.Println("Question: ", q.question)
		answer, err := answerQuestion()
		if err != nil {
			log.Fatalf("failed to answer question: %v", err)
		}
		if answer == q.answer {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect!")
		}
	}

}

//func main() {
//	nums := []int{10, 20, 30}
//	for index, value := range nums {
//		fmt.Println(index, value)
//	}
//}
