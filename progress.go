package quiz

// Progress tracks the amount of correct questions
type Progress struct {
	total    int
	answered int
	correct  int
	done     bool
}

// AddAnswer modifies tracking state
func (p *Progress) AddAnswer(correct bool) {
	p.answered++
	if correct {
		p.correct++
	}
}

// IsDone tells if all questions are answered
func (p *Progress) IsDone() bool {
	return p.answered == p.total
}

// NextIndex returns index of next question
func (p *Progress) NextIndex() int {
	return p.answered
}

// Correct number of questions answered
func (p *Progress) Correct() int {
	return p.correct
}
