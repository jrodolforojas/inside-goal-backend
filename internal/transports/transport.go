package transports

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/jrodolforojas/inside-goal-backend/internal/service"
	"github.com/gorilla/mux"
	httptransport "github.com/go-kit/kit/transport/http"
)

type errorer interface {
	error() error
}

// MakeHTTPHandler mounts all of the service endpoints into an http.Handler.
func MakeHTTPHandler(ctx context.Context, s *service.Feed) http.Handler {
	router := mux.NewRouter()
	endpoints := service.MakeServerEndpoints(s)

	// GET		/orders/{id}	get an order by ID Order
	router.Methods("GET").Path("/news").Handler(httptransport.NewServer(
		endpoints.GetNews,
		decodeGetNewsRequest,
		encodeResponse,
	))

	return router
}

func decodeGetNewsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req service.GetNewsRequest
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
