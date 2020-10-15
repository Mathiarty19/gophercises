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
	a := getAnswer(q)
	c , t:= quizResults(q, a)
	fmt.Printf("You got %d correct out of %d", c, t)


}
//Test contains questions and answers from csv
type Test struct {
	question string
	answer string
}

//takes the argument csv file and formats it in a slice of structs
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
	return questions
}

//tallys the amount of questions asked and gotten correct
func quizResults(q []*Test, a  []string) (int, int) {
	correct := 0
	total := 0
	for i := 0; i < 12; i++{
		switch{
		case q[i].answer == a[i]:
			correct++
			total++
		case q[i].answer != "":
			total++
		default:
			break
		}
	}
	return correct, total
}

//prints questions to the terminal and records the user response with loop break at 30 seconds
func getAnswer(q []*Test) []string  {
	answers := make([]string, 12)
	fmt.Println("Test starting you have thirty seconds")
	start := time.Now()
	for i := 0; i < 12; i++{
		if time.Since(start) > time.Second * 30 {
			fmt.Println("Times Up!")
			break
		} else {
		fmt.Println(q[i].question)
		fmt.Scanf("%s", &answers[i])
		}
	}
	return answers
}
