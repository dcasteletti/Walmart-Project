package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"proyect/config"
	"proyect/entity"
	"strings"
)

type Repository struct {
}

var collection = config.GetCollection("products")
var ctx = context.Background()

func GetAllProducts() (*entity.Products, error) {
	products := &entity.Products{}

	request, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	for request.Next(ctx) {
		product := &entity.Product{}
		err = request.Decode(product)
		if err != nil {
			return nil, err
		}

		products.Product = append(products.Product, product)
	}

	return products, nil
}

func GetProductByID(value string) (*entity.Products, error) {
	products := &entity.Products{}
	filter := bson.D{{"id", value}}

	cursor := collection.FindOne(context.Background(), filter)
	if cursor == nil {
		return nil, errors.New("")
	}

	var product *entity.Product

	_ = cursor.Decode(&product)

	products.Product = append(products.Product, product)

	return products, nil
}

func GetProducts(value string) (*entity.Products, error) {
	isPalindrome := isPalindrome(value)
	products := &entity.Products{}

	// regex := fmt.Sprintf("/%s/", value)

	filter := bson.D{
		{Key: "$or", Value: bson.A{
			bson.D{{Key: "description", Value: value}},
			bson.D{{Key: "brand", Value: primitive.Regex{Pattern: value}}},
		}}}
	sortedBy := options.Find().SetSort(bson.D{{"id", 1}})

	cursor, err := collection.Find(context.Background(), filter, sortedBy)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var product *entity.Product

		if err := cursor.Decode(&product); err != nil {
			continue
		}

		if isPalindrome {
			halfPercentDiscount(product)
		}

		products.Product = append(products.Product, product)
	}

	return products, nil
}

func isPalindrome(word string) bool {
	word = strings.ToUpper(word)

	for i, letter := range word {
		finalWord := len(word) - 1 - i

		if uint8(letter) != word[finalWord] {
			return false
		}
	}

	return true
}

func halfPercentDiscount(product *entity.Product) {
	product.Price = product.Price / 2
}
