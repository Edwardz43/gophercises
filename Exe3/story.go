package main

import (
	"encoding/json"
	"io/ioutil"
)

func getStory() *Story {
	var story *Story

	s, err := ioutil.ReadFile("./gopher.json")

	check(err)

	err = json.Unmarshal(s, &story)

	return story
}

func mapStory(url string, s *Story) interface{} {
	switch url {
	case "/":
		return s.Intro

	case "/new-york":
		return s.NewYork

	case "/debate":
		return s.Debate

	case "/sean-kelly":
		return s.SeanKelly

	case "/mark-bates":
		return s.MarkBates

	case "/denver":
		return s.Denver

	case "/home":
		return s.Home
	default:
		return nil
	}
}

// func readStory(story Story, next chan []byte) {
// 	for k, v := range story.Intro.Story {

// 		// if k == len(story.Intro.Story) {
// 		// fmt.Printf("%s(%s)", v, "NewYork/Denver")
// 		// } else {
// 		fmt.Printf("%s(%d)\n", v, k)
// 		// }

// 		reader := bufio.NewReader(os.Stdin)
// 		in, err := reader.ReadBytes('\n')
// 		check(err)
// 		next <- in
// 	}
// }
