package server

import (
	h "github.com/Markovk1n/go-kit-tutor1/pkg/handler"
	"github.com/Markovk1n/go-kit-tutor1/pkg/services"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func NewHTTPServer(endpoints h.Endpoints) *http.ServeMux {
	router := http.NewServeMux()

	// создадим простой middleware
	// он будет устанавливать для всех запросов,
	// зарегистрированных через него, тип ответа "application/json"

	handle := func(pattern string, handler http.Handler) {
		router.Handle(pattern, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Add("Content-Type", "application/json; charset=utf-8")
			handler.ServeHTTP(writer, request)
		}))
	}

	handle("/status", httptransport.NewServer(
		endpoints.StatusEndpoint,
		services.DecodeStatusRequest,
		services.EncodeResponse,
	),
	)

	handle(
		"/get",
		httptransport.NewServer(
			endpoints.GetEndpoint,
			services.DecodeGetRequest,
			services.EncodeResponse,
		),
	)

	handle(
		"/validate",
		httptransport.NewServer(
			endpoints.ValidateEndpoint,
			services.DecodeValidateRequest,
			services.EncodeResponse,
		),
	)

	return router
}
