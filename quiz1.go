package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	q := csvOpen()
	a := getAnswer()
	quiz(q, a)

}

func csvOpen() []string {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	r := csv.NewReader(f)
	var record []string
	for {
		record, err = r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
	return record
}

func quiz(q []string, a string) bool {
	// if answer and their given answer match return true ect
}

func getAnswer() string {

	fmt.Println(resp)
	return resp
}
