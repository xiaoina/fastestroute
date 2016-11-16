package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
)

func makeRouteEndpoint(svc RouteService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(routeRequest)
		v, err := svc.Route(req.S)
		if err != nil {
			return routeResponse{v, err.Error()}, nil
		}
		return routeResponse{v, ""}, nil
	}
}

func decodeRouteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request routeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeRouteResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response routeResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func encodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

type routeRequest struct {
	S []string `json:"s"`
}

type routeResponse struct {
	V   []GoogleMapsResponse `json:"v"`
	Err string               `json:"err,omitempty"`
}
