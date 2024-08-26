package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sushistack/link.stack/configs"
	"github.com/sushistack/link.stack/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func TestNewMongoDBClient(t *testing.T) {
	utils.InitProjectRoot()
	config := configs.LoadConfig(nil)
	client := NewMongoDBClient(config.Datasource)

	assert.NotNil(t, client)
}

func TestInsertDocument(t *testing.T) {
	utils.InitProjectRoot()
	config := configs.LoadConfig(nil)
	client := NewMongoDBClient(config.Datasource)

	result, err := client.InsertDocument("testdb", "testcollection", bson.M{"name": "test"})
	assert.NoError(t, err)
	assert.NotNil(t, result)
	client.DeleteDocument("testdb", "testcollection", bson.M{"name": "test"})
}

func TestFindDocument(t *testing.T) {
	utils.InitProjectRoot()
	config := configs.LoadConfig(nil)
	client := NewMongoDBClient(config.Datasource)

	client.InsertDocument("testdb", "testcollection", bson.M{"name": "test"})
	result, err := client.FindDocument("testdb", "testcollection", bson.M{"name": "test"})
	assert.NoError(t, err)
	assert.NotNil(t, result)
	client.DeleteDocument("testdb", "testcollection", bson.M{"name": "test"})
}

func TestUpdateDocument(t *testing.T) {
	utils.InitProjectRoot()
	config := configs.LoadConfig(nil)
	client := NewMongoDBClient(config.Datasource)

	client.InsertDocument("testdb", "testcollection", bson.M{"name": "test"})
	result, err := client.UpdateDocument("testdb", "testcollection", bson.M{"name": "test"}, bson.M{"$set": bson.M{"name": "updated"}})
	assert.NoError(t, err)
	assert.NotNil(t, result)
	client.DeleteDocument("testdb", "testcollection", bson.M{"name": "updated"})
}

func TestDeleteDocument(t *testing.T) {
	utils.InitProjectRoot()
	config := configs.LoadConfig(nil)
	client := NewMongoDBClient(config.Datasource)

	client.InsertDocument("testdb", "testcollection", bson.M{"name": "test"})
	result, err := client.DeleteDocument("testdb", "testcollection", bson.M{"name": "test"})
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
