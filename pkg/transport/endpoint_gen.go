//+build !swipe

// Code generated by Swipe v2.0.0-beta6. DO NOT EDIT.

package transport

import (
	"context"

	"github.com/fesyunoff/phone-book/pkg/service"
	"github.com/fesyunoff/phone-book/pkg/service/dto"
	"github.com/go-kit/kit/endpoint"
)

func makeServiceAddEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		result, err := s.Add(ctx, req.Task)
		if err != nil {
			return nil, err
		}
		return result, nil
	}

}

func makeServiceDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		result, err := s.Delete(ctx, req.Task)
		if err != nil {
			return nil, err
		}
		return result, nil
	}

}

func makeServiceGetEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		out, msg, err := s.Get(ctx, req.Task)
		if err != nil {
			return nil, err
		}
		return GetResponse{Out: out, Msg: msg}, nil
	}

}

func makeServiceUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		result, err := s.Update(ctx, req.Task)
		if err != nil {
			return nil, err
		}
		return result, nil
	}

}

type ServiceEndpointSet struct {
	AddEndpoint    endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
	GetEndpoint    endpoint.Endpoint
	UpdateEndpoint endpoint.Endpoint
}

func MakeServiceEndpointSet(svc service.Service) ServiceEndpointSet {
	return ServiceEndpointSet{
		AddEndpoint:    makeServiceAddEndpoint(svc),
		DeleteEndpoint: makeServiceDeleteEndpoint(svc),
		GetEndpoint:    makeServiceGetEndpoint(svc),
		UpdateEndpoint: makeServiceUpdateEndpoint(svc),
	}
}

type AddRequest struct {
	Task dto.Entry `json:"task"`
}
type DeleteRequest struct {
	Task dto.Entry `json:"task"`
}
type GetRequest struct {
	Task dto.Entry `json:"task"`
}
type GetResponse struct {
	Out []*dto.Entry `json:"out"`
	Msg string       `json:"msg"`
}
type UpdateRequest struct {
	Task dto.Entry `json:"task"`
}