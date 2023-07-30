// https://courses.calhoun.io/lessons/les_goph_01

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var cvsPtr = flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer' (default 'problems.csv')")
var limPtr = flag.Int("limit", 30, "the time limit for the quiz in seconds (default 30)")
var shufPtr = flag.Bool("shuffle", false, "shuffle questios randomly")
var scoreQ = 0
var totalQ = 0

func readCSV(file_name string) (data [][]string) {
	f, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err = csvReader.ReadAll()

	totalQ = len(data)

	if *shufPtr {
		rand.Shuffle(len(data), func(i, j int) {
			data[i], data[j] = data[j], data[i]
		})
	}

	if err != nil {
		log.Fatal(err)
	}

	return
}

func read_quizes(quizes [][]string) {

	var inp string

	for i, d := range quizes {

		fmt.Printf("Problem #%d: %s \n", i+1, d[0])
		fmt.Scan(&inp)

		inp = strings.ToLower(strings.TrimSpace(inp))

		if inp == d[1] {
			scoreQ++
		}
	}

}

func time_quizes(t int) {

	duration := time.Duration(t) * time.Second

	time.AfterFunc(duration, func() {
		summary_quiz()
		os.Exit(0)
	})

}

func summary_quiz() {
	fmt.Printf("You scored %d out of %d", scoreQ, totalQ)

}

func main() {

	rand.New(rand.NewSource(time.Now().UnixNano()))
	flag.Parse()

	quizes := readCSV(*cvsPtr)

	fmt.Println("To start Quiz press Enter.")
	fmt.Scanln()

	time_quizes(*limPtr)
	read_quizes(quizes)
	summary_quiz()

}
