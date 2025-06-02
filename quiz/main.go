package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
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

func runTimer(numSeconds int) {
	timer := time.NewTimer(time.Duration(numSeconds))
	<-timer.C

	fmt.Println("Timer has expired")
}

func main() {
	file, err := os.Open("quiz/problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var responses []string

	// Create channels for timer and input
	timeUp := make(chan bool)
	inputChan := make(chan string)

	// Start the timer
	go func() {
		time.Sleep(30 * time.Second)
		timeUp <- true
	}()

	// Main game loop
	for r, err := reader.Read(); err == nil; r, err = reader.Read() {
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(r[0])

		// Start goroutine to get user input
		go func() {
			var answer string
			fmt.Scanln(&answer)
			inputChan <- answer
		}()

		// Wait for either timer or input
		select {
		case <-timeUp:
			fmt.Println("\nTime's up!")
			fmt.Println("Final score:", responses)
			return
		case answer := <-inputChan:
			if answer == r[1] {
				responses = append(responses, "correct")
			} else {
				responses = append(responses, "incorrect")
			}
		}
	}

	fmt.Println(responses)
}
