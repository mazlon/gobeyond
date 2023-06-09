package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mazlon/gobeyond/internal/models"
)

type QuestionResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func AskQuestions(db sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		query := `insert into questions (question, date_create) 
		values ($1, NOW()) 
		RETURNING id`
		var jsonData models.Questions
		defer c.Request.Body.Close()
		decoder := json.NewDecoder(c.Request.Body)
		err := decoder.Decode(&jsonData)
		if err != nil {
			log.Println("error while unmarshaling")
		}
		questionOnFly, _ := json.Marshal(jsonData.Question)
		dbRes, err := db.Query(query, questionOnFly)
		if err != nil {
			log.Println("error while inserting data")
			log.Println(err)
		}
		fmt.Println("database result is: ", dbRes)
		resp := QuestionResponse{}
		resp.Message = "We added your question, no Guaranty for you to get your answer back"
		resp.Status = 200
		respJson, err := json.Marshal(resp)
		if err != nil {
			log.Fatal("unable to unmarshaling")
		}
		c.Writer.Write(respJson)
	}
}
