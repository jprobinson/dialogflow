package dialogflow

import (
	"context"
	"net/http"

	"google.golang.org/appengine"

	"github.com/NYTimes/marvin"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type service struct {
	google     map[string]GoogleActionHandler
	middleware endpoint.Middleware
}

func Run(google GoogleActionServer, middleware endpoint.Middleware) {
	marvin.Init(service{google: google.Actions()})
	appengine.Main()
}

func (s service) Options() []httptransport.ServerOption {
	return []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			httptransport.EncodeJSONResponse(ctx, w, err)
		}),
	}
}

func (s service) RouterOptions() []marvin.RouterOption {
	return nil
}

func (s service) HTTPMiddleware(h http.Handler) http.Handler {
	return h
}

func (s service) Middleware(ep endpoint.Endpoint) endpoint.Endpoint {
	if s.Middleware != nil {
		return s.middleware(ep)
	}
	return ep
}

func (s service) JSONEndpoints() map[string]map[string]marvin.HTTPEndpoint {
	return map[string]map[string]marvin.HTTPEndpoint{
		"/google": {
			"POST": {
				Endpoint: s.postGoogle,
				Decoder:  decodeGoogle,
			},
		},
	}
}

var errBadRequest = marvin.NewJSONStatusResponse(map[string]string{
	"error": "bad request"}, http.StatusBadRequest)
