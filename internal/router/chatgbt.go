package router

import (
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

// func AskQuestions(db sql.DB) func(*gin.Context) {
// 	// We are using closure here to pass db to handler which normally doesn't accept anything but gin context
// 	return func(c *gin.Context) {
// 		query := `insert into questions (question, date_create)
// 		values ($1, NOW())
// 		RETURNING id`
// 		var jsonData models.Questions
// 		defer c.Request.Body.Close()
// 		decoder := json.NewDecoder(c.Request.Body)
// 		err := decoder.Decode(&jsonData)
// 		if err != nil {
// 			log.Println("error while unmarshaling")
// 		}
// 		questionOnFly, _ := json.Marshal(jsonData.Question)
// 		dbRes, err := db.Query(query, questionOnFly)
// 		if err != nil {
// 			log.Println("error while inserting data")
// 			log.Println(err)
// 		}
// 		fmt.Println("database result is: ", dbRes)
// 		resp := QuestionResponse{}
// 		resp.Message = "We added your question, no Guaranty for you to get your answer back"
// 		resp.Status = 200
// 		respJson, err := json.Marshal(resp)
// 		if err != nil {
// 			log.Fatal("unable to unmarshaling")
// 		}
// 		c.Writer.Write(respJson)
// 	}
// }

func (gbt *GbtServer) AskQuestions(c *gin.Context) {
	query := `insert into questions (question, date_create, user_id) 
		values ($1, NOW(),$2) 
		RETURNING id`
	var jsonData models.Questions
	defer c.Request.Body.Close()
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&jsonData)
	if err != nil {
		log.Println("error while unmarshaling")
	}
	questionOnFly, _ := json.Marshal(jsonData.Question)
	dbRes, err := gbt.dbConnection.Query(query, questionOnFly, c.Keys[`userID`])
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
