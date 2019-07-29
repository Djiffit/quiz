package quiz

// Question struct
type Question struct {
	Answer     string
	Question   string
	UserAnswer string
}

// AnswerQuestion returns if answer was correct
func (q *Question) AnswerQuestion(answer string) bool {
	q.UserAnswer = answer
	return q.Answer == answer
}
