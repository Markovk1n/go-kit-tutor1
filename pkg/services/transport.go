package services

import (
	"context"
	"encoding/json"
	"github.com/Markovk1n/go-kit-tutor1/pkg/models"
	"net/http"
)

func DecodeGetRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return models.GetRequest{}, nil
}

func DecodeValidateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req models.ValidateRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeStatusRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return models.StatusRequest{}, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
