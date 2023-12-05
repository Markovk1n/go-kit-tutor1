package handler

import (
	"context"
	"errors"
	m "github.com/Markovk1n/go-kit-tutor1/pkg/models"
	s "github.com/Markovk1n/go-kit-tutor1/pkg/services"

	"github.com/go-kit/kit/endpoint"
)

var (
	errUnexpected = errors.New("unexpected error")
)

func MakeStatusEndpoint(srv s.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := srv.Status(ctx)
		if err != nil {
			return m.StatusResponse{res}, err
		}

		return m.StatusResponse{res}, nil
	}
}

func MakeGetEndpoint(srv s.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := srv.Get(ctx)
		if err != nil {
			return m.GetResponse{res, err}, err
		}

		return m.GetResponse{Date: res}, nil
	}
}

func MakeValidateEndpoint(srv s.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(m.ValidateRequest)
		if !ok {
			return m.ValidateResponse{Err: errUnexpected}, errUnexpected
		}

		res, err := srv.Validate(ctx, req.Date)
		if err != nil {
			return m.ValidateResponse{res, err}, nil
		}

		return m.ValidateResponse{Valid: res}, nil
	}
}

//////

type Endpoints struct {
	StatusEndpoint   endpoint.Endpoint
	GetEndpoint      endpoint.Endpoint
	ValidateEndpoint endpoint.Endpoint
}

func (e Endpoints) Get(ctx context.Context) (string, error) {
	req := m.GetRequest{}

	resp, err := e.GetEndpoint(ctx, req)
	if err != nil {
		return "", err
	}

	getResp, ok := resp.(m.GetResponse)
	if !ok {
		return "", errUnexpected
	}

	if getResp.Err != nil {
		return "", getResp.Err
	}

	return getResp.Date, nil
}

func (e Endpoints) Status(ctx context.Context) (string, error) {
	req := m.StatusRequest{}

	resp, err := e.StatusEndpoint(ctx, req)
	if err != nil {
		return "", err
	}

	statusResp, ok := resp.(m.StatusResponse)
	if !ok {
		return "", errUnexpected
	}

	return statusResp.Status, nil
}

func (e Endpoints) Validate(ctx context.Context, date string) (bool, error) {
	req := m.ValidateRequest{Date: date}

	resp, err := e.ValidateEndpoint(ctx, req)
	if err != nil {
		return false, err
	}

	validateResp, ok := resp.(m.ValidateResponse)
	if !ok {
		return false, errUnexpected
	}

	if validateResp.Err != nil {
		return false, validateResp.Err
	}

	return validateResp.Valid, nil
}
