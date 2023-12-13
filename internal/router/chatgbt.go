package router

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mazlon/gobeyond/internal/messaging"
	"github.com/mazlon/gobeyond/internal/models"
)

type QuestionResponse struct {
	Status    int       `json:"status"`
	Message   string    `json:"message"`
	MessageID uuid.UUID `json:"message_id"`
}

func (gbt *GbtServer) AskQuestions(c *gin.Context) {
	query := `
	INSERT INTO questions (question, date_create, user_id) 
	VALUES ($1, NOW(), $2) 
	RETURNING id
`
	var jsonData models.Questions
	var questionRow models.Questions
	defer c.Request.Body.Close()
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&jsonData)
	if err != nil {
		log.Println("error while unmarshaling:", err)
		// Handle the error appropriately (e.g., return an error response).
		return
	}
	// Check if the database connection is established.
	if gbt.dbConnection == nil {
		log.Println("database connection is not established")
		// Handle the error appropriately (e.g., return an error response).
		return
	}
	// Use a prepared statement to execute the query.
	stmt, err := gbt.dbConnection.Prepare(query)
	if err != nil {
		log.Println("error while preparing query:", err)
		// Handle the error appropriately (e.g., return an error response).
		return
	}
	defer stmt.Close()
	// Execute the query and retrieve the ID using Scan.
	err = stmt.QueryRow(jsonData.Question, c.Keys["userID"]).Scan(&questionRow.ID)
	if err != nil {
		log.Println("error while inserting data:", err)
		// Handle the error appropriately (e.g., return an error response).
		return
	}
	fmt.Println("Inserted ID:", questionRow.ID)
	questionRow.Question = jsonData.Question
	resp := QuestionResponse{}
	resp.Message = "We added your question, no Guaranty for you to get your answer back"
	resp.Status = 200
	resp.MessageID = questionRow.ID
	respJson, err := json.Marshal(resp)
	if err != nil {
		log.Println("unable to unmarshaling")
	}
	queue := gbt.queue
	log.Println("Enqueuing the question...")
	err = messaging.EnqueuingQuestions(questionRow, queue)
	if err != nil {
		log.Println("Unable to enqueue the question", err)
	}

	c.Writer.Write(respJson)
}
