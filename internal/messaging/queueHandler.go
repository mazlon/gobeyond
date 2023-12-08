package messaging

import (
	"context"
	"log"

	"golang.org/x/sync/errgroup"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vgarvardt/gue/v5"
)

const (
	printerQueue    = "name_printer"
	questionQueue   = "questions_queue"
	jobTypePrinter  = "PrintName"
	jobTypeQuestion = "questionHandler"
)

type printNameArgs struct {
	Name string
}

type questionText struct {
	Content string
}

func Queue(ctx context.Context, connectionPool *pgxpool.Pool) (gc *gue.Client, err error) {
	gc, err = NewMessagingClient(ctx, connectionPool)
	if err != nil {
		log.Println("Error in creating a messaging client")
		return nil, err
	}
	wm := gue.WorkMap{
		jobTypePrinter:  printName,
		jobTypeQuestion: askEnqueuedQuestionsFromApi,
	}

	// create a pool w/ 2 workers
	workers, err := gue.NewWorkerPool(gc, wm, 2)
	if err != nil {
		log.Println(err)
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
			log.Println(err)
		}
		return err
	})

	// args, err := json.Marshal(printNameArgs{Name: "vgarvardt"})
	// if err != nil {
	// 	log.Println(err)
	// }

	// j := &gue.Job{
	// 	Type:  jobTypePrinter,
	// 	Queue: printerQueue,
	// 	Args:  args,
	// }
	// if err := gc.Enqueue(ctx, j); err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }

	// j = &gue.Job{
	// 	Type:  jobTypePrinter,
	// 	Queue: printerQueue,
	// 	RunAt: time.Now().UTC().Add(30 * time.Second), // delay 30 seconds
	// 	Args:  args,
	// }
	// if err := gc.Enqueue(ctx, j); err != nil {
	// 	log.Fatal(err)
	// 	return nil, err
	// }

	// time.Sleep(30 * time.Second) // wait for while
	// // send shutdown signal to worker
	// if err := g.Wait(); err != nil {
	// 	log.Fatal(err)
	// 	return nil, err
	// }
	return gc, nil
}
