package services

import (
	"context"
	"time"
)

type Service interface {
	Status(ctx context.Context) (string, error)
	Get(ctx context.Context) (string, error)
	Validate(ctx context.Context, date string) (bool, error)
}

type dateService struct{}

func NewService() Service {
	return dateService{}
}

func (s dateService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

func (s dateService) Get(ctx context.Context) (string, error) {
	now := time.Now()
	return now.Format("02.01.2006"), nil
}
func (s dateService) Validate(ctx context.Context, date string) (bool, error) {
	_, err := time.Parse("02.01.2006", date)
	if err != nil {
		return false, err
	}

	return true, nil
}
