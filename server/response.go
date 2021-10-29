package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type generalResponse struct {
	Errors  []*ErrorResponse `json:"errors"`
	Success bool             `json:"success"`
	State   interface{}      `json:"state"`
}

type ErrOption func(*ErrorResponse)

func (r *Response) JSON(code int, data interface{}) {
	r.writeJSON(code, nil, data)
}

func (r *Response) Error(code int, msg string, opts ...ErrOption) {
	err := &ErrorResponse{Code: code, Message: msg}

	r.writeJSON(code, []*ErrorResponse{err}, nil)
}

func (r *Response) writeJSON(code int, errors []*ErrorResponse, state interface{}) {
	r.writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	response := &generalResponse{Errors: errors, State: state}
	b, err := json.Marshal(response)
	if err != nil {
		r.writer.WriteHeader(http.StatusInternalServerError)
		r.writer.Write([]byte(fmt.Sprintf("unexpected error: %v", err)))
	}

	r.writer.WriteHeader(code)
	r.writer.Write(b)
}
