package ai

import (
	"github.com/mazlon/gobeyond/internal/models"
)

type QuestionAsker interface {
	AskQuestion(question models.Questions) (models.Answers, error)
}
