package quiz

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

// RunQuiz runs the quiz game
func RunQuiz(file *os.File, reader *bufio.Scanner, timer, totalTimer int) {
	quiz := CreateNewQuiz(file)

	fmt.Printf("You have %d seconds for all questions and %d seconds for each question", totalTimer, timer)
	promptNextQuestion("Are you ready to play the game?", reader)

	questionsCh := make(chan bool)
	timerCh := make(chan bool)

	go askQuestions(&quiz, reader, timer, questionsCh)
	go runTimer(timerCh, totalTimer)

	select {
	case <-questionsCh:
		fmt.Printf("Congratulations you managed to answer all questions!!")
	case <-timerCh:
		fmt.Printf("Too bad! You ran out of time.")
	}

	printSummary(quiz.questions)
	fmt.Println(quiz.progress.Correct(), " / ", quiz.progress.total, " Correct!")
	promptNextQuestion("Press enter to exit the program ...", reader)
}

func askQuestions(quiz *Quiz, reader *bufio.Scanner, timer int, ch chan bool) {
	for i := 0; i < len(quiz.GetQuestions()); i++ {

		answerCh := make(chan string)
		timerCh := make(chan bool)

		go getAnswer(quiz, reader, answerCh)
		go runTimer(timerCh, timer)

		select {
		case ans := <-answerCh:

			fmt.Print(ans)
			quiz.Answer(ans)
		case <-timerCh:
			quiz.Answer("")
		}

		if !quiz.progress.IsDone() {
			promptNextQuestion("Ready for next question?", reader)
		}

	}

	ch <- true
}

func printSummary(questions []Question) {
	fmt.Printf("\n\n\n -----------------  \n\n")
	for i, question := range questions {
		fmt.Printf("Question #%d was %q. You answered: %q, the correct answer was %q \n\n", i+1, question.Question, question.UserAnswer, question.Answer)
	}
}

func promptNextQuestion(message string, r *bufio.Scanner) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println(message, " Please press enter :^)")
	readLine(r)
}

func runTimer(ch chan bool, duration int) {
	time.Sleep(time.Duration(duration) * time.Second)
	ch <- true
}

func getAnswer(quiz *Quiz, reader *bufio.Scanner, ch chan string) {
	q, err := quiz.GetQuestion()

	if err != nil {
		log.Fatalf("Error getting next question %v", err)
	}

	fmt.Println("Next question: ", q)
	ch <- readLine(reader)
}

func readLine(reader *bufio.Scanner) string {
	reader.Scan()
	return reader.Text()
}
