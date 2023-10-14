package healing

import (
	"context"
	"log/slog"
	"slices"
	"strings"
)

// QueueClient hypotetical queue client struct.
type QueueClient struct {
	values []string
	logger *slog.Logger
}

func NewQueueClient(values []string, logger *slog.Logger) *QueueClient {
	newQueueClient := new(QueueClient)

	newQueueClient.values = values
	newQueueClient.setLogger(logger)

	newQueueClient.logger.Info(
		"initializing queue client",
		slog.String("values", strings.Join(values, ",")),
	)

	return newQueueClient
}

func (q *QueueClient) RetrieveMessages(ctx context.Context) []string {
	q.logger.Info("retrieving messages", slog.String("values", strings.Join(q.values, ",")))
	return slices.Clone(q.values)
}

func (q *QueueClient) setLogger(logger *slog.Logger) {
	q.logger = logger.With(slog.String("component", "QueueClient"))
}
