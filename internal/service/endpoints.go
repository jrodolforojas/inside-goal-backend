package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/jrodolforojas/inside-goal-backend/internal/models"
)

type Endpoints struct {
	GetNews      endpoint.Endpoint
	GetProviders endpoint.Endpoint
}

func MakeServerEndpoints(service *Feed) Endpoints {
	return Endpoints{
		GetNews:      makeGetNewsEndpoint(service),
		GetProviders: makeGetProvidersEndpoint(service),
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

type GetNewsRequest struct {
}

// GetNewsResponse struct
type GetNewsResponse struct {
	News []models.Notice `json:"news,omitempty"`
	Err  error           `json:"err,omitempty"`
}

func (r GetNewsResponse) Error() error { return r.Err }

func makeGetProvidersEndpoint(service *Feed) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		providers, err := service.GetProviders(ctx)

		if err != nil {
			return GetProvidersResponse{Providers: nil, Err: err}, err
		}
		return GetProvidersResponse{Providers: providers, Err: err}, nil
	}
}

type GetProvidersRequest struct {
}

// GetProvidersResponse struct
type GetProvidersResponse struct {
	Providers []models.Provider `json:"providers,omitempty"`
	Err       error             `json:"err,omitempty"`
}

func (r GetProvidersResponse) Error() error { return r.Err }
