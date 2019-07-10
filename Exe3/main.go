package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	var story Story

	s, err := ioutil.ReadFile("./gopher.json")

	check(err)

	json.Unmarshal(s, &story)

	for {
		next := make(chan []byte, len(story.Intro.Story))
		go readStory(story, next)
		select {
		case <-next:
			break
		}
	}
}

func readStory(story Story, next chan []byte) {
	for k, v := range story.Intro.Story {

		// if k == len(story.Intro.Story) {
		// fmt.Printf("%s(%s)", v, "NewYork/Denver")
		// } else {
		fmt.Printf("%s(%d)\n", v, k)
		// }

		reader := bufio.NewReader(os.Stdin)
		in, err := reader.ReadBytes('\n')
		check(err)
		next <- in
	}
}
