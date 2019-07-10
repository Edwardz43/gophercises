package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

// Profile ...
var profile struct {
	user  string
	score int8
}

var total int8

var limit time.Duration

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	preProcess()

	// Read data from csv
	dat, err := ioutil.ReadFile("../problems.csv")

	check(err)

	r := csv.NewReader(strings.NewReader(string(dat)))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s = ", record[0])

		input := make(chan string, 1)

		go getInput(input)

		select {
		case i := <-input:
			if i == record[1] {

				profile.score++
			}
		case <-time.After(limit):
			fmt.Println("\ntime up !")
			break
		}
		total++
	}

	fmt.Printf("Done !\n%s,You scored %d out of %d\n", profile.user, profile.score, total)
}

func preProcess() {
	flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")

	l := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	flag.Parse()

	if *l == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	limit = time.Second * time.Duration(*l)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Your name, pls")

	u, err := reader.ReadString('\n')
	check(err)

	// remove \n from lines
	profile.user = strings.TrimSuffix(string(u), "\n")

	fmt.Printf("Hello %s, ready to test !\n", profile.user)

	time.Sleep(time.Second * 1)
}

func getQuizs() {

}

func getInput(input chan string) {
	for {
		reader := bufio.NewReader(os.Stdin)

		in, err := reader.ReadBytes('\n')
		check(err)

		input <- strings.TrimSuffix(string(in), "\n")
	}
}
