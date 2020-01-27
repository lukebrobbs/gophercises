package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func handleErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func handleInput(item []string) bool {
	fmt.Print(item[0] + ": ")

	var input string
	_, err := fmt.Scanln(&input)
	handleErr(err)
	if input != item[1] {
		return false
	}
	return true
}

func main() {

	csvFile := flag.String("f", "./problems.csv", "file path for csv file")
	timeout := flag.Int("t", 30, "Amount of time(in seconds) allowed to complete the quiz")
	flag.Parse()

	dat, err := ioutil.ReadFile(*csvFile)
	handleErr(err)

	r := csv.NewReader(strings.NewReader(string(dat)))
	records, err := r.ReadAll()
	handleErr(err)

	var score int
	timer := time.NewTimer(time.Duration(*timeout) * time.Second)

	for index, item := range records {
		go func() {
			<-timer.C
			fmt.Printf("out of time, you answered %d questions out of %d, and got %d correct\n", index, len(records), score)
			os.Exit(1)
		}()
		if handleInput(item) == true {
			score++
		}
	}
	fmt.Printf("you got %d out of %d correct! \n", score, len(records))
}
