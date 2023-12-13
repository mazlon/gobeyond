package messaging

import (
	"context"
	"log"

	"golang.org/x/sync/errgroup"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vgarvardt/gue/v5"
)

const (
	questionQueue   = "questions_queue"
	jobTypeQuestion = "questionHandler"
)

func Queue(ctx context.Context, connectionPool *pgxpool.Pool) (gc *gue.Client, err error) {
	gc, err = NewMessagingClient(ctx, connectionPool)
	if err != nil {
		log.Println("Error in creating a messaging client")
		return nil, err
	}
	wm := gue.WorkMap{
		jobTypeQuestion: askEnqueuedQuestionsFromApi,
	}

	// create a pool w/ 2 workers
	workers, err := gue.NewWorkerPool(gc, wm, 2, gue.WithPoolQueue(questionQueue), gue.WithPoolHooksJobDone(finishedJobsLog))
	if err != nil {
		log.Println("Error creating working pool", err)
		return nil, err
	}

	// work jobs in goroutine
	g, gctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		err := workers.Run(gctx)
		if err != nil {
			// In a real-world applications, use a better way to shut down
			// application on unrecoverable error. E.g. fx.Shutdowner from
			// go.uber.org/fx module.
			log.Println("Error in worker pool", err)
		}
		return err
	})

	return gc, nil
}
