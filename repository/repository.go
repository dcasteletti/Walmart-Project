package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetProducts(value string) (*entity.Products, error) {
	if len(value) < 3 {
		return nil, errors.New("value must be longer than 3")
	}

	isPalindrome := isPalindrome(value)
	products := &entity.Products{}

	filter := bson.D{
		{"$or", bson.A{
			bson.D{{"id", value}},
			bson.D{{Key: "description", Value: primitive.Regex{Pattern: value}}},
			bson.D{{Key: "brand", Value: primitive.Regex{Pattern: value}}},
		}}}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var product *entity.Product

		if err := cursor.Decode(&product); err != nil {
			fmt.Println(err)
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
