package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	oracle := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		oracle <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second * 10)
			go prophecy("", answers)
		}
	}()

	go func() {
		for question := range questions {
			go answer(question, answers)
		}
	}()

	go func() {
		for answer := range answers {
			fmt.Printf(answer)
			time.Sleep(time.Duration(20+rand.Intn(10)) * time.Millisecond)
		}
	}()

	return questions
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(20+rand.Intn(10)) * time.Second)

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
		"You should make another question.",
		"I think you can decide.",
		"You are funny",
	}
	answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))] + "\n" + prompt
}

func answer(question string, answer chan<- string) {
	fmt.Println("Answer is coming...")
	time.Sleep(time.Millisecond * 20)
	randAnswers := []string{
		"The answer is blowing in the wind..",
		"Follow your heart!",
		"You know better than I do.",
		"Yeah you should know.",
		"Nothing matters.",
	}
	answer <- randAnswers[rand.Intn(len(randAnswers))] + "\n" + prompt
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
