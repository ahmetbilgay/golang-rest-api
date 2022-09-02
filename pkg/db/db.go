package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"golang-rest-api/pkg/utils"
)

func ConnectDB() {
	dbUri := utils.LoadDotEnv("DB_URI")
	col := utils.ConnectCollection(dbUri, "blgy-app", "users")
	data := bson.D{{"title", "Record of a Shriveled Datum"}, {"text", "No bytes, no problem. Just insert a document, in MongoDB"}}
	result, err := col.InsertOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
