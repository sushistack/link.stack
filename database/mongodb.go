package database

import (
	"context"
	"github.com/sushistack/link.stack/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDBClient struct {
	client *mongo.Client
}

func NewMongoDBClient(config utils.Config) *MongoDBClient {
	uri := config.Datasource.URI
	connectionPoolMinSize := config.Datasource.ConnectionPool.MinSize
	connectionPoolMaxSize := config.Datasource.ConnectionPool.MaxSize
	// connectionPoolMaxIdle := config.Datasource.ConnectionPool.MaxIdle

	clientOptions := options.Client().
		ApplyURI(uri).
		SetMaxPoolSize(connectionPoolMinSize).
		SetMinPoolSize(connectionPoolMaxSize).
		SetMaxConnIdleTime(5 * time.Minute)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		utils.Logger.Error("Can not connect db.", err)
		return nil
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		utils.Logger.Error("Can not received pong.", err)
		return nil
	}

	return &MongoDBClient{
		client: client,
	}
}

func (dbClient *MongoDBClient) InsertDocument(databaseName, collectionName string, document interface{}) (*mongo.InsertOneResult, error) {
	collection := dbClient.client.Database(databaseName).Collection(collectionName)
	return collection.InsertOne(context.Background(), document)
}

func (dbClient *MongoDBClient) FindDocument(databaseName, collectionName string, filter interface{}) (bson.M, error) {
	collection := dbClient.client.Database(databaseName).Collection(collectionName)
	var result bson.M
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (dbClient *MongoDBClient) UpdateDocument(databaseName, collectionName string, filter, update interface{}) (*mongo.UpdateResult, error) {
	collection := dbClient.client.Database(databaseName).Collection(collectionName)
	return collection.UpdateOne(context.Background(), filter, update)
}

func (dbClient *MongoDBClient) DeleteDocument(databaseName, collectionName string, filter interface{}) (*mongo.DeleteResult, error) {
	collection := dbClient.client.Database(databaseName).Collection(collectionName)
	return collection.DeleteOne(context.Background(), filter)
}
