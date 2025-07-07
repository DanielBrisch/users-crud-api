package logger

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoHook struct {
	Collection *mongo.Collection
}

func (h *MongoHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *MongoHook) Fire(entry *logrus.Entry) error {
	logDoc := map[string]any{
		"level":     entry.Level.String(),
		"message":   entry.Message,
		"timestamp": entry.Time,
		"fields":    entry.Data,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := h.Collection.InsertOne(ctx, logDoc)
	return err
}
