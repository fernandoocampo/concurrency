package healing

import (
	"fmt"
	"log/slog"
	"slices"
)

type repo struct {
	logger  *slog.Logger
	storage []string
}

func NewRepository(logger *slog.Logger) *repo {
	newRepo := &repo{
		storage: make([]string, 0),
	}

	newRepo.setLogger(logger)

	return newRepo
}

func (r *repo) Save(record string) error {
	r.logger.Info("saving new record", slog.String("record", record))

	if slices.Contains(r.storage, record) {
		r.logger.Error("saving a new record",
			slog.String("record", record),
			slog.String("error", "record already exists"))

		return fmt.Errorf("record already exists")
	}

	r.storage = append(r.storage, record)

	r.logger.Info("record was saved successfully", slog.String("record", record))

	return nil
}

func (r *repo) Read(index int) string {
	r.logger.Info("reading a record", slog.Int("record_id", index))

	if index < 0 || index >= len(r.storage) {
		return ""
	}

	return r.storage[index]
}

func (r *repo) setLogger(logger *slog.Logger) {
	r.logger = logger.With(slog.String("component", "repo"))
}
