package transports

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/jrodolforojas/inside-goal-backend/internal/middleware"
	"github.com/jrodolforojas/inside-goal-backend/internal/service"
)

type errorer interface {
	error() error
}

// MakeHTTPHandler mounts all of the service endpoints into an http.Handler.
func MakeHTTPHandler(ctx context.Context, s *service.Feed) http.Handler {
	router := mux.NewRouter()
	endpoints := service.MakeServerEndpoints(s)

	subRouter := router.PathPrefix("/api").Subrouter()
	router = router.PathPrefix("/").Subrouter()

	corsMethods := []string{http.MethodOptions, http.MethodGet}
	router.Use(middleware.CORSPolicies(corsMethods))
	subRouter.Use(middleware.CORSPolicies(corsMethods))

	router.Methods(http.MethodGet).Path("/news").Handler(httptransport.NewServer(
		endpoints.GetNews,
		decodeGetNewsRequest,
		encodeResponse,
	))

	router.Methods(http.MethodGet).Path("/providers").Handler(httptransport.NewServer(
		endpoints.GetProviders,
		decodeGetProvidersRequest,
		encodeResponse,
	))

	return router
}

func decodeGetNewsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req service.GetNewsRequest
	return req, nil
}

func decodeGetProvidersRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req service.GetProvidersRequest
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		log.Println("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

var (
	// ErrInconsistentIDs inconsistent IDs error returns int error.
	errInconsistentIDs = errors.New("inconsistent IDs")
	// ErrAlreadyExists already exists error returns int error.
	errAlreadyExists = errors.New("already exists")
	// ErrNotFound not found error returns int error.
	errNotFound = errors.New("not found")
)

func codeFrom(err error) int {
	switch err {
	case errNotFound:
		return http.StatusNotFound
	case errAlreadyExists, errInconsistentIDs:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
