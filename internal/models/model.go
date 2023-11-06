// all the database structures
// We can switch to service layer and use interfaces instead of calling DB iteractions directly within handlers
package models

type Questions struct {
	Question string `json:"question"`
}
