package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var mission int
var guess int

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	fmt.Println("_________________________")
	fmt.Println("Welcome to the mini-game!")
	var timer time.Duration
	fmt.Println("Select time limit (1-3):")
	fmt.Println("1) 60sec. - beginner")
	fmt.Println("2) 30sec. - skilled")
	fmt.Println("3) 15sec. - cybersport")
	fmt.Fscan(os.Stdin, &timer)
	if timer == 1 {
		timer = 60
	} else if timer == 2 {
		timer = 30
	} else if timer == 3 {
		timer = 15
	} else {
		fmt.Println("Invalid value!")
		os.Exit(0)
	}
	f := func() {
		fmt.Println("__________________________")
		fmt.Println("TIMER IS OVER!!! TRY AGAIN")
		os.Exit(0)
	}
	time.AfterFunc(time.Second*timer, f)
	rand.Seed(time.Now().UnixNano())
	mission = randomInt(0, 100)
	//fmt.Println(mission)
	GuessTheNumber()
}

func GuessTheNumber() {

	fmt.Println("Input your namber(0-100):")
	fmt.Fscan(os.Stdin, &guess)
	if guess > mission {
		fmt.Println("Take less..")
		fmt.Println("")
		GuessTheNumber()
	} else if guess < mission {
		fmt.Println("Take bigger..")
		fmt.Println("")
		GuessTheNumber()
	} else if guess == mission {
		fmt.Println("__________________________")
		fmt.Println("You guessed right! WINNER!")
		os.Exit(0)
	}

}
