package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
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

	wordPtr := flag.String("f", "./problems.csv", "file path for csv file")
	flag.Parse()

	dat, err := ioutil.ReadFile(*wordPtr)
	handleErr(err)

	r := csv.NewReader(strings.NewReader(string(dat)))
	records, err := r.ReadAll()
	handleErr(err)

	var score int

	for _, item := range records {
		if handleInput(item) == true {
			score++
		}
	}
	fmt.Printf("you got %v out of %v correct! \n", score, len(records))
}
