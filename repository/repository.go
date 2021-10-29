package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"proyect/config"
	"proyect/entity"
	"strconv"
	"strings"
)

type Repository struct {
}

var ctx = context.Background()

func GetCollection() (*mongo.Collection, error) {
	collection, err := config.GetCollection("products")
	if err != nil {
		return nil, fmt.Errorf("could not connect to database:%s", err)
	}

	return collection, err
}

func GetAllProducts() (*entity.Products, error) {
	collection, err := GetCollection()
	if err != nil {
		return nil, err
	}

	products := &entity.Products{}

	request, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	for request.Next(ctx) {
		product := &entity.Product{}
		err = request.Decode(product)
		if err != nil {
			continue
		}

		products.Product = append(products.Product, product)
	}

	return products, nil
}

func GetProductByID(id string) (*entity.Products, error) {
	collection, err := GetCollection()
	if err != nil {
		return nil, err
	}

	idInt, err := strconv.Atoi(id)
	products := &entity.Products{}
	filter := bson.D{{Key: "id", Value: idInt}}

	cursor, err := collection.Find(context.TODO(), filter, options.Find().SetLimit(1))
	for cursor.Next(ctx) {
		var product *entity.Product

		if err := cursor.Decode(&product); err != nil {
			continue
		}

		products.Product = append(products.Product, product)
	}

	return products, nil
}

func GetProducts(value string) (*entity.Products, error) {
	collection, err := GetCollection()
	if err != nil {
		return nil, err
	}

	products := &entity.Products{}

	filter := bson.D{
		{Key: "$or", Value: bson.A{
			bson.D{{Key: "description", Value: primitive.Regex{Pattern: value}}},
			bson.D{{Key: "brand", Value: primitive.Regex{Pattern: value}}},
		}}}
	sortedBy := options.Find().SetSort(bson.D{{"id", 1}})

	cursor, err := collection.Find(context.Background(), filter, sortedBy)
	if err != nil || cursor == nil {
		return nil, fmt.Errorf("error in query %s", err)
	}

	for cursor.Next(ctx) {
		var product *entity.Product

		if err := cursor.Decode(&product); err != nil {
			continue
		}

		if IsPalindrome(value) {
			product.Price = HalfPercentDiscount(product.Price)
		}

		products.Product = append(products.Product, product)
	}

	return products, nil
}

func IsPalindrome(word string) bool {
	word = strings.ToUpper(word)

	for i, letter := range word {
		finalWord := len(word) - 1 - i

		if uint8(letter) != word[finalWord] {
			return false
		}
	}

	return true
}

func HalfPercentDiscount(price int) int {
	return price / 2
}
