package main

import (
	"time"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	q := csvOpen()
	// a := getAnswer(q)
	// quizResults(q, a)
	getAnswer(q)

}
type Test struct {
	question string
	answer string
}

func csvOpen() []*Test {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	r := csv.NewReader(f)
	questions := make([]*Test, 12)
	var i int = 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		questions[i] = &Test{question: record[0], answer: record[1]}
		i++
	}
	// fmt.Println(questions)
	return questions
}

// func quizResults(q []string, a string) bool {
// 	// if answer and their given answer match return true ect
// }

func getAnswer(questions []*Test) []string  {
	answers := make([]string, 12)
	fmt.Println("Test starting you have thirty seconds")
	start := time.Now()
	for i := 0; i < 12; i++{
		var answer string
		if time.Since(start) > time.Second * 30 {
			fmt.Println("Times Up!")
			break
		} else {
		fmt.Println(questions[i].question)
		fmt.Scanf("%s", &answer)
		answers = append(answers, answer)
		}
	}
	return answers
}
