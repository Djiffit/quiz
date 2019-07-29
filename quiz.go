package quiz

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"
)

// Quiz struct contains all questions and progress for current quiz
type Quiz struct {
	questions []Question
	progress  Progress
	mu        sync.Mutex
}

// GetQuestions returns questions
func (q *Quiz) GetQuestions() []Question {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.questions
}

//GetProgress resturns progress
func (q *Quiz) GetProgress() Progress {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.progress
}

// CreateNewQuiz returns a new quiz based on a file
func CreateNewQuiz(file *os.File) Quiz {
	var questions []Question

	err := json.NewDecoder(file).Decode(&questions)

	if err != nil {
		log.Fatalf("Failed to read the given questionnaire %v", err)
	}

	return Quiz{
		questions,
		Progress{total: len(questions)},
		sync.Mutex{},
	}
}

// GetQuestion returns next question
func (q *Quiz) GetQuestion() (string, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if !q.progress.IsDone() {
		return q.questions[q.progress.NextIndex()].Question, nil
	}

	return "", errors.New("No more questions")
}

// Answer the next question
func (q *Quiz) Answer(answer string) (corr bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if !q.progress.IsDone() {
		corr = q.questions[q.progress.NextIndex()].AnswerQuestion(answer)
		q.progress.AddAnswer(corr)
	}
	return
}
