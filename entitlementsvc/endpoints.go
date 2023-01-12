package entitlementsvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	GetEntitlementEndpoint endpoint.Endpoint
}
type getEntitlementRequest struct {
	ID string
}

type getEntitlementResponse struct {
	Entitlement Entitlement `json:"entitlement,omitempty"`
	Err         error       `json:"err,omitempty"`
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetEntitlementEndpoint: MakeGetEntitlementEndpoint(s),
	}
}

// GetProfile implements Service. Primarily useful in a client.
func (e Endpoints) GetEntitlement(ctx context.Context, id string) (Entitlement, error) {
	request := getEntitlementRequest{ID: id}
	response, err := e.GetEntitlementEndpoint(ctx, request)
	if err != nil {
		return Entitlement{}, err
	}
	resp := response.(getEntitlementResponse)

	return resp.Entitlement, resp.Err
}

// MakeGetEntitlementEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakeGetEntitlementEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getEntitlementRequest)
		p, e := s.GetEntitlement(ctx, req.ID)
		return getEntitlementResponse{Entitlement: p, Err: e}, nil
	}
}
