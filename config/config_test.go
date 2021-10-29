package config_test

import (
	"context"
	"github.com/mongo-go/testdb"
	"go.mongodb.org/mongo-driver/mongo"
	"proyect/config"
	"testing"
	"time"
)

func TestGetCollection(t *testing.T) {
	coll, _ := config.GetCollection("products")

	defer coll.Database()
}

func setup(t *testing.T) *mongo.Collection {
	var testDb *testdb.TestDB

	if testDb == nil {
		testDb = testdb.NewTestDB("mongodb://localhost", "local", time.Duration(2) * time.Second)

		err := testDb.Connect()
		if err != nil {
			t.Fatal(err)
		}
	}

	coll, err := testDb.CreateRandomCollection(testdb.NoIndexes)
	if err != nil {
		t.Fatal(err)
	}

	return coll
}

func Test1(t *testing.T) {
	coll := setup(t)
	defer coll.Drop(context.Background())
}