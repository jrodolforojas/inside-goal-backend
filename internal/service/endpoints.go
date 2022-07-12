package service

import (
	"context"

	"github.com/jrodolforojas/inside-goal-backend/internal/models"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetNews endpoint.Endpoint
}

func MakeServerEndpoints(service *Feed) Endpoints {
	return Endpoints{
		GetNews: makeGetNewsEndpoint(service),
	}
}

func makeGetNewsEndpoint(service *Feed) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		news, err := service.GetNews(ctx)

		if err != nil {
			return GetNewsResponse{News: nil, Err: err}, err
		}
		return GetNewsResponse{News: news, Err: err}, nil
	}
}

type GetNewsRequest struct{
	
}

// GetNewsResponse struct
type GetNewsResponse struct {
	News []models.Notice `json:"news,omitempty"`
	Err  error           `json:"err,omitempty"`
}

func (r GetNewsResponse) Error() error { return r.Err }
