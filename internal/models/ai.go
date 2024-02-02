package models

import (
	"github.com/google/uuid"
)

type Questions struct {
	Question string `json:"question"`
	ID       uuid.UUID
}

type Answers struct {
	Answer     string    `json:"answer"`
	QuestionId uuid.UUID `json:"question_id"`
}
