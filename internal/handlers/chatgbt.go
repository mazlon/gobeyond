package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type Questions struct {
	Question string `json:"question"`
}

type QuestionResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func AskQuestions(c *gin.Context) {
	fmt.Println(c.Request.Body)
	defer c.Request.Body.Close()
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("error while unmarshaling")
	}
	fmt.Println(jsonData)
	resp := QuestionResponse{}
	resp.Message = "We added your question, no Guaranty for you to ge t your answer back"
	resp.Status = 200
	respJson, err := json.Marshal(resp)
	if err != nil {
		log.Fatal("unable to unmarshaling")
	}
	c.Writer.Write(respJson)
}
