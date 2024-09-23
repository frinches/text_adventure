package main

import "fmt"

type Scene struct {
	id          string
	description string
	choices     []Choice
	isEnding    bool
}

type Choice struct {
	text        string
	nextSceneID string
}

func (s *Scene) Display() {
	fmt.Println(s.description)
	fmt.Println()
	
	if !s.isEnding {
		fmt.Println("What will you do?")
		for i, choice := range s.choices {
			fmt.Printf("%d. %s\n", i+1, choice.text)
		}
	}
}