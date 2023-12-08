package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vgarvardt/gue/v5"
	"github.com/vgarvardt/gue/v5/adapter/pgxv5"
)

type fetchDBArgs struct {
	UserID int
}

type questions struct {
	question string
}

func printName(ctx context.Context, j *gue.Job) error {
	var args printNameArgs
	if err := json.Unmarshal(j.Args, &args); err != nil {
		return err
	}
	fmt.Printf("Hello %s!\n", args.Name)
	return nil
}

func finishedJobsLog(ctx context.Context, j *gue.Job, err error) {
	if err != nil {
		return
	}

	j.Tx().Exec(
		ctx,
		"INSERT INTO finished_jobs_log (queue, type, run_at) VALUES ($1, $2, now())",
		j.Queue,
		j.Type,
	)
}

// func fetchAndPrintFromDB(ctx context.Context, j *gue.Job) error {
// 	var args fetchDBArgs
// 	if err := json.Unmarshal(j.Args, &args); err != nil {
// 		return err
// 	}

// 	var question questions
// 	err := pgxpool.QueryRow(ctx, "SELECT question FROM users WHERE id=$1", args.UserID).Scan(&question.question)
// 	if err != nil {
// 		return err
// 	}

// 	// Print the data obtained
// 	fmt.Printf("Question: %d, ID: %s\n", question.question, args.UserID)

// 	return nil
// }

func EnqueuingQuestions(question string, qClient *gue.Client) error {
	questionContent, err := json.Marshal(questionText{Content: question})
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
	var question questionText
	if err := json.Unmarshal(j.Args, &question); err != nil {
		fmt.Print("Error while unmarshaling the question before sending to API")
		return err
	}
	fmt.Print(question)
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
