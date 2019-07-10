package main

import (
	"bufio"
	"encoding/csv"
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	fmt.Println("Your name, pls")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		profile.user = scanner.Text()
		fmt.Printf("Hello %s, ready to test !\n", scanner.Text())
	}

	time.Sleep(time.Second * 2)

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

		// fmt.Println("make some input : ")
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			in := scanner.Text()

			if in == record[1] {
				// fmt.Println("bingo")
				profile.score++
			}
			total++
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}
		}
	}

	fmt.Printf("Done !\n%s,You scored %d out of %d\n", profile.user, profile.score, total)

}
