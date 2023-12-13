// all the database structures
// We can switch to service layer and use interfaces instead of calling DB iteractions directly within handlers
package models

import (
	"github.com/google/uuid"
)
type Questions struct {
	Question string `json:"question"`
	ID       uuid.UUID
}
