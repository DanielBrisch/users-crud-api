package logger

import (
	"context"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitLogger() *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal("Error on connect to MongoDB:", err)
	}

	collection := client.Database("logsdb").Collection("applogs")

	log.AddHook(&MongoHook{Collection: collection})

	return log
}
