package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/Djiffit/quiz"
)

func main() {
	src := flag.String("src", "quiz.json", "file name")
	timer := flag.Int("timer", 3, "Time to answer a question")
	totalTimer := flag.Int("totalTimer", 30, "Time to answer all questions")

	flag.Parse()

	fmt.Print(*src)

	json, _ := os.Open(*src)

	quiz.RunQuiz(json, bufio.NewScanner(os.Stdin), *timer, *totalTimer)
}
