package server

import "net/http"

type Response struct {
	writer http.ResponseWriter
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{
		writer: w,

	}
}

