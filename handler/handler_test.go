package handler

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProducts(t *testing.T) {

	tests := []struct {
		host     string
		expected string
	}{
		{host: "local", expected: "" },
	}

	for _, test := range tests {
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r.Host = test.host
		id := r.Header.Get("id")
		assert.Equal(t, test.expected, id)
	}

}
