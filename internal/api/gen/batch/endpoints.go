// Code generated by goa v3.8.1, DO NOT EDIT.
//
// batch endpoints
//
// Command:
// $ goa-v3.8.1 gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package batch

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "batch" service endpoints.
type Endpoints struct {
	Submit goa.Endpoint
	Status goa.Endpoint
	Hints  goa.Endpoint
}

// NewEndpoints wraps the methods of the "batch" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Submit: NewSubmitEndpoint(s),
		Status: NewStatusEndpoint(s),
		Hints:  NewHintsEndpoint(s),
	}
}

// Use applies the given middleware to all the "batch" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Submit = m(e.Submit)
	e.Status = m(e.Status)
	e.Hints = m(e.Hints)
}

// NewSubmitEndpoint returns an endpoint function that calls the method
// "submit" of service "batch".
func NewSubmitEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SubmitPayload)
		return s.Submit(ctx, p)
	}
}

// NewStatusEndpoint returns an endpoint function that calls the method
// "status" of service "batch".
func NewStatusEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Status(ctx)
	}
}

// NewHintsEndpoint returns an endpoint function that calls the method "hints"
// of service "batch".
func NewHintsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Hints(ctx)
	}
}
