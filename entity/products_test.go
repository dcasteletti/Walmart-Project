package entity_test

import (
	"github.com/stretchr/testify/assert"
	"proyect/entity"
	"testing"
)

func TestProductEntity(t *testing.T) {
	t.Parallel()

	product := &entity.Product{
		Id: 1,
		Description: "TV",
		Brand: "",
		Price: 990,
		Image: "www.walmart.com",
	}
	
	assert.NotNil(t, product)
}