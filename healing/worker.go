package healing

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"
)

const (
	pulseInterval   = 20 * time.Millisecond
	timeoutInterval = 1000 * time.Millisecond
)

// Worker defines logic to process messages.
type Worker struct {
	queueClient     *QueueClient
	repository      *repo
	logger          *slog.Logger
	pulseInterval   time.Duration
	timeoutInterval time.Duration
}

func NewWorker(queueClient *QueueClient, repository *repo, logger *slog.Logger) *Worker {
	logger.Info("initializing worker")

	newWorker := Worker{
		queueClient:     queueClient,
		repository:      repository,
		logger:          newWorkerLogger(logger),
		pulseInterval:   pulseInterval,
		timeoutInterval: timeoutInterval,
	}

	newWorker.logger.Info("new worker has been initialized")

	return &newWorker
}

func (w *Worker) ProcessMessages(ctx context.Context) error {
	w.logger.Info("starting to process messages")
	timeoutSignal := time.NewTicker(w.timeoutInterval)
	eventStreamCtx, eventStreamCancel := context.WithCancel(ctx)
	eventStream, eventHeartbeat := w.readEvents(eventStreamCtx)
	for {
		select {
		case <-ctx.Done():
			w.logger.Info("context was cancelled while processing messages", "err", ctx.Err())
			eventStreamCancel()
			return fmt.Errorf("processing messages: %w", ctx.Err())
		case <-eventHeartbeat:
			w.logger.Info("resetting the timeout signal")
			timeoutSignal.Reset(w.timeoutInterval)
		case <-timeoutSignal.C:
			w.logger.Info("event stream unhealthy... restarting")
			w.logger.Info("canceling event reading stream")
			eventStreamCancel()
			w.logger.Info("restarting event reading stream")
			eventStreamCtx, eventStreamCancel = context.WithCancel(ctx)
			eventStream, eventHeartbeat = w.readEvents(eventStreamCtx)
			continue
		case event, ok := <-eventStream:
			if !ok {
				w.logger.Info("event message stream was closed")
				eventStreamCancel()
				return nil
			}

			err := w.save(event)
			if err != nil {
				w.logger.Error("saving event", "err", err)
			}
		}
	}
}

func (w *Worker) readEvents(ctx context.Context) (<-chan string, <-chan any) {
	w.logger.Info("creating goroutine to read events")

	readLogger := w.logger.With(slog.String("goroutine", "event_reader"))

	pulse := time.NewTicker(w.pulseInterval)
	eventStream := make(chan string)
	heartbeat := make(chan any)
	sendPulseFunc := func(ctx context.Context, heartbeat chan<- any, aLogger *slog.Logger) {
		readLogger.Info("ready to send a pulse")
		select {
		case <-ctx.Done():
			readLogger.Info("context was cancelled while there was a pulse", "err", ctx.Err())
			return
		case heartbeat <- true:
			readLogger.Info("sending heartbeat pulse")
		default:
		}
	}
	go func() {
		defer close(eventStream)
		defer close(heartbeat)

		readLogger.Info("starting to read events")

		for {
			select {
			case <-ctx.Done():
				readLogger.Info("context was cancelled while retrieving messages", "err", ctx.Err())
				return
			case <-pulse.C:
				sendPulseFunc(ctx, heartbeat, readLogger)
			default:
				readLogger.Info("reading new messages")
				newMessages := w.queueClient.RetrieveMessages(ctx)
				readLogger.Info("reading new messages", slog.Int("messages", len(newMessages)))
				for _, message := range newMessages {
					readLogger.Info("reading new message", slog.String("message", message))
					select {
					case <-ctx.Done():
						w.logger.Info("context was cancelled while reading a message", "err", ctx.Err())
						return
					case eventStream <- message:
					}
				}
			}
		}
	}()
	return eventStream, heartbeat
}

func (w *Worker) save(record string) error {
	w.logger.Info("saving a new record", slog.String("record", record))
	err := w.repository.Save(strings.ToUpper(record))
	if err != nil {
		w.logger.Error(
			"saving new record",
			slog.String("record", record),
			"err", err,
		)

		return fmt.Errorf("unable to save record: %w", err)
	}

	w.logger.Info("new record was created", slog.String("record", record))

	return nil
}

func newWorkerLogger(logger *slog.Logger) *slog.Logger {
	return logger.With(slog.String("component", "worker"))
}
