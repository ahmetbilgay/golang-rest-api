package utils

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func LoadDotEnv(envName string) string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	envData := os.Getenv(envName)
	if envData == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	return envData
}

func ConnectDB(uri string) {
	_, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
}

func ConnectCollection(uri string, dbName string, colName string) *mongo.Collection {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	coll := client.Database(dbName).Collection(colName)
	return coll
}
