package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mazlon/gobeyond/internal/ai"
	"github.com/mazlon/gobeyond/internal/models"
	"github.com/vgarvardt/gue/v5"
	"github.com/vgarvardt/gue/v5/adapter/pgxv5"
)

func finishedJobsLog(ctx context.Context, j *gue.Job, err error) {
	if err != nil {
		return
	}

	_, err = j.Tx().Exec(
		ctx,
		"INSERT INTO finished_jobs_log (queue_id, queue_body, run_at) VALUES ($1, $2, now())",
		j.ID,
		string(j.Args[:]),
	)
	if err != nil {
		// Used Fatal here because if this function isn't working, the jobs will run infinitely
		log.Fatal(err)
	}
}

func EnqueuingQuestions(question models.Questions, qClient *gue.Client) error {
	questionContent, err := json.Marshal(question)
	if err != nil {
		fmt.Print("An error while marshaling question before Enqueue!")
		return err
	}
	job := &gue.Job{
		Type:  jobTypeQuestion,
		Queue: questionQueue,
		// RunAt: time.Now().UTC().Add(30 * time.Second), // delay 30 seconds
		Args: questionContent,
	}
	err = qClient.Enqueue(context.Background(), job)
	if err != nil {
		fmt.Print("An error while Queuing the Job Question")
		return err
	}
	return nil
}

func askEnqueuedQuestionsFromApi(ctx context.Context, j *gue.Job) error {
	var question models.Questions
	if err := json.Unmarshal(j.Args, &question); err != nil {
		log.Println("Error while unmarshaling the question before sending to API:", err)
		return err
	}
	log.Println("Question: ", question.Question, "ID: ", question.ID)
	gptResponse, err := ai.AskFromChatGptSingleQuestion(&question)
	if err != nil {
		log.Println("An error while askign the question from chatGPT: ", err)
		return err
	}
	log.Println("message: ", gptResponse.Choices[0].Message.Content, " qId: ", gptResponse.QuestionId)
	_, err = j.Tx().Exec(
		ctx,
		"INSERT INTO answers (answer, question_id, date_create) VALUES ($1, $2, now())",
		gptResponse.Choices[0].Message.Content,
		gptResponse.QuestionId,
	)
	if err != nil {
		log.Println("An error while inserting the answer: ", err)
		return err
	}
	return nil
}

func NewMessagingClient(ctx context.Context, connectionPool *pgxpool.Pool) (*gue.Client, error) {
	poolAdapter := pgxv5.NewConnPool(connectionPool)
	err := connectionPool.Ping(ctx)
	if err != nil {
		return nil, err
	}
	gc, err := gue.NewClient(poolAdapter)
	if err != nil {
		log.Println("Error while calling gue new client")
	}
	return gc, nil
}
