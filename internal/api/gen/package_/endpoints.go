// Code generated by goa v3.8.1, DO NOT EDIT.
//
// package endpoints
//
// Command:
// $ goa-v3.8.1 gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package package_

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "package" service endpoints.
type Endpoints struct {
	Monitor             goa.Endpoint
	List                goa.Endpoint
	Show                goa.Endpoint
	Delete              goa.Endpoint
	Cancel              goa.Endpoint
	Retry               goa.Endpoint
	Bulk                goa.Endpoint
	BulkStatus          goa.Endpoint
	PreservationActions goa.Endpoint
	Confirm             goa.Endpoint
	Reject              goa.Endpoint
	Move                goa.Endpoint
	MoveStatus          goa.Endpoint
}

// MonitorEndpointInput holds both the payload and the server stream of the
// "monitor" method.
type MonitorEndpointInput struct {
	// Stream is the server stream used by the "monitor" method to send data.
	Stream MonitorServerStream
}

// NewEndpoints wraps the methods of the "package" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Monitor:             NewMonitorEndpoint(s),
		List:                NewListEndpoint(s),
		Show:                NewShowEndpoint(s),
		Delete:              NewDeleteEndpoint(s),
		Cancel:              NewCancelEndpoint(s),
		Retry:               NewRetryEndpoint(s),
		Bulk:                NewBulkEndpoint(s),
		BulkStatus:          NewBulkStatusEndpoint(s),
		PreservationActions: NewPreservationActionsEndpoint(s),
		Confirm:             NewConfirmEndpoint(s),
		Reject:              NewRejectEndpoint(s),
		Move:                NewMoveEndpoint(s),
		MoveStatus:          NewMoveStatusEndpoint(s),
	}
}

// Use applies the given middleware to all the "package" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Monitor = m(e.Monitor)
	e.List = m(e.List)
	e.Show = m(e.Show)
	e.Delete = m(e.Delete)
	e.Cancel = m(e.Cancel)
	e.Retry = m(e.Retry)
	e.Bulk = m(e.Bulk)
	e.BulkStatus = m(e.BulkStatus)
	e.PreservationActions = m(e.PreservationActions)
	e.Confirm = m(e.Confirm)
	e.Reject = m(e.Reject)
	e.Move = m(e.Move)
	e.MoveStatus = m(e.MoveStatus)
}

// NewMonitorEndpoint returns an endpoint function that calls the method
// "monitor" of service "package".
func NewMonitorEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		ep := req.(*MonitorEndpointInput)
		return nil, s.Monitor(ctx, ep.Stream)
	}
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "package".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListPayload)
		return s.List(ctx, p)
	}
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "package".
func NewShowEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ShowPayload)
		res, err := s.Show(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedEnduroStoredPackage(res, "default")
		return vres, nil
	}
}

// NewDeleteEndpoint returns an endpoint function that calls the method
// "delete" of service "package".
func NewDeleteEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeletePayload)
		return nil, s.Delete(ctx, p)
	}
}

// NewCancelEndpoint returns an endpoint function that calls the method
// "cancel" of service "package".
func NewCancelEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CancelPayload)
		return nil, s.Cancel(ctx, p)
	}
}

// NewRetryEndpoint returns an endpoint function that calls the method "retry"
// of service "package".
func NewRetryEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RetryPayload)
		return nil, s.Retry(ctx, p)
	}
}

// NewBulkEndpoint returns an endpoint function that calls the method "bulk" of
// service "package".
func NewBulkEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*BulkPayload)
		return s.Bulk(ctx, p)
	}
}

// NewBulkStatusEndpoint returns an endpoint function that calls the method
// "bulk_status" of service "package".
func NewBulkStatusEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.BulkStatus(ctx)
	}
}

// NewPreservationActionsEndpoint returns an endpoint function that calls the
// method "preservation-actions" of service "package".
func NewPreservationActionsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*PreservationActionsPayload)
		res, err := s.PreservationActions(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedEnduroPackagePreservationActions(res, "default")
		return vres, nil
	}
}

// NewConfirmEndpoint returns an endpoint function that calls the method
// "confirm" of service "package".
func NewConfirmEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ConfirmPayload)
		return nil, s.Confirm(ctx, p)
	}
}

// NewRejectEndpoint returns an endpoint function that calls the method
// "reject" of service "package".
func NewRejectEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RejectPayload)
		return nil, s.Reject(ctx, p)
	}
}

// NewMoveEndpoint returns an endpoint function that calls the method "move" of
// service "package".
func NewMoveEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*MovePayload)
		return nil, s.Move(ctx, p)
	}
}

// NewMoveStatusEndpoint returns an endpoint function that calls the method
// "move_status" of service "package".
func NewMoveStatusEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*MoveStatusPayload)
		return s.MoveStatus(ctx, p)
	}
}
