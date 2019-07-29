package quiz_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/Djiffit/quiz"
)

func TestQuiz(t *testing.T) {
	initialData := `[
		{"answer": "10", "question": "5 + 5"}]`

	t.Run("The quiz is initialized correctly and can be answered", func(t *testing.T) {
		quizSource, remove := createTempFile(t, initialData)

		defer remove()

		quiz := quiz.CreateNewQuiz(quizSource)
		quiz.Answer("10")

		assertProgress(t, quiz.GetProgress(), 1, 1)

	})

	t.Run("Can't answer to more questions than there are", func(t *testing.T) {
		quizSource, remove := createTempFile(t, initialData)

		defer remove()

		quiz := quiz.CreateNewQuiz(quizSource)
		quiz.Answer("10")
		quiz.Answer("10")
		quiz.Answer("10")

		assertProgress(t, quiz.GetProgress(), 1, 1)

	})
}

func assertProgress(t *testing.T, progress quiz.Progress, answered, correct int) {
	t.Helper()

	if progress.NextIndex() != answered {
		t.Errorf("%v, invalid number of questions answered, expected %d, got %d", progress, progress.NextIndex(), answered)
	}

	if progress.Correct() != correct {
		t.Errorf("%v, invalid number of correctly answered questions %d, expected %d", progress, progress.Correct(), correct)
	}
}

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	if initialData == "" {
		initialData = "[]"
	}

	tmpfile, err := ioutil.TempFile("", "questionnaire_test.json")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}
	tmpfile.Seek(0, 0)

	return tmpfile, removeFile
}
