// Code generated by goa v3.8.1, DO NOT EDIT.
//
// storage HTTP server
//
// Command:
// $ goa-v3.8.1 gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package server

import (
	"context"
	"net/http"

	storage "github.com/artefactual-sdps/enduro/internal/api/gen/storage"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the storage service endpoint HTTP handlers.
type Server struct {
	Mounts       []*MountPoint
	Submit       http.Handler
	Update       http.Handler
	Download     http.Handler
	Locations    http.Handler
	AddLocation  http.Handler
	Move         http.Handler
	MoveStatus   http.Handler
	Reject       http.Handler
	Show         http.Handler
	ShowLocation http.Handler
	CORS         http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the storage service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *storage.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Submit", "POST", "/storage/package/{aip_id}/submit"},
			{"Update", "POST", "/storage/package/{aip_id}/update"},
			{"Download", "GET", "/storage/package/{aip_id}/download"},
			{"Locations", "GET", "/storage/location"},
			{"AddLocation", "POST", "/storage/location"},
			{"Move", "POST", "/storage/package/{aip_id}/store"},
			{"MoveStatus", "GET", "/storage/package/{aip_id}/store"},
			{"Reject", "POST", "/storage/package/{aip_id}/reject"},
			{"Show", "GET", "/storage/package/{aip_id}"},
			{"ShowLocation", "GET", "/storage/location/{uuid}"},
			{"CORS", "OPTIONS", "/storage/package/{aip_id}/submit"},
			{"CORS", "OPTIONS", "/storage/package/{aip_id}/update"},
			{"CORS", "OPTIONS", "/storage/package/{aip_id}/download"},
			{"CORS", "OPTIONS", "/storage/location"},
			{"CORS", "OPTIONS", "/storage/package/{aip_id}/store"},
			{"CORS", "OPTIONS", "/storage/package/{aip_id}/reject"},
			{"CORS", "OPTIONS", "/storage/package/{aip_id}"},
			{"CORS", "OPTIONS", "/storage/location/{uuid}"},
		},
		Submit:       NewSubmitHandler(e.Submit, mux, decoder, encoder, errhandler, formatter),
		Update:       NewUpdateHandler(e.Update, mux, decoder, encoder, errhandler, formatter),
		Download:     NewDownloadHandler(e.Download, mux, decoder, encoder, errhandler, formatter),
		Locations:    NewLocationsHandler(e.Locations, mux, decoder, encoder, errhandler, formatter),
		AddLocation:  NewAddLocationHandler(e.AddLocation, mux, decoder, encoder, errhandler, formatter),
		Move:         NewMoveHandler(e.Move, mux, decoder, encoder, errhandler, formatter),
		MoveStatus:   NewMoveStatusHandler(e.MoveStatus, mux, decoder, encoder, errhandler, formatter),
		Reject:       NewRejectHandler(e.Reject, mux, decoder, encoder, errhandler, formatter),
		Show:         NewShowHandler(e.Show, mux, decoder, encoder, errhandler, formatter),
		ShowLocation: NewShowLocationHandler(e.ShowLocation, mux, decoder, encoder, errhandler, formatter),
		CORS:         NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "storage" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Submit = m(s.Submit)
	s.Update = m(s.Update)
	s.Download = m(s.Download)
	s.Locations = m(s.Locations)
	s.AddLocation = m(s.AddLocation)
	s.Move = m(s.Move)
	s.MoveStatus = m(s.MoveStatus)
	s.Reject = m(s.Reject)
	s.Show = m(s.Show)
	s.ShowLocation = m(s.ShowLocation)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the storage endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountSubmitHandler(mux, h.Submit)
	MountUpdateHandler(mux, h.Update)
	MountDownloadHandler(mux, h.Download)
	MountLocationsHandler(mux, h.Locations)
	MountAddLocationHandler(mux, h.AddLocation)
	MountMoveHandler(mux, h.Move)
	MountMoveStatusHandler(mux, h.MoveStatus)
	MountRejectHandler(mux, h.Reject)
	MountShowHandler(mux, h.Show)
	MountShowLocationHandler(mux, h.ShowLocation)
	MountCORSHandler(mux, h.CORS)
}

// Mount configures the mux to serve the storage endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountSubmitHandler configures the mux to serve the "storage" service
// "submit" endpoint.
func MountSubmitHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleStorageOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/storage/package/{aip_id}/submit", f)
}

// NewSubmitHandler creates a HTTP handler which loads the HTTP request and
// calls the "storage" service "submit" endpoint.
func NewSubmitHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeSubmitRequest(mux, decoder)
		encodeResponse = EncodeSubmitResponse(encoder)
		encodeError    = EncodeSubmitError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "submit")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountUpdateHandler configures the mux to serve the "storage" service
// "update" endpoint.
func MountUpdateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleStorageOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/storage/package/{aip_id}/update", f)
}

// NewUpdateHandler creates a HTTP handler which loads the HTTP request and
// calls the "storage" service "update" endpoint.
func NewUpdateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateRequest(mux, decoder)
		encodeResponse = EncodeUpdateResponse(encoder)
		encodeError    = EncodeUpdateError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "update")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountDownloadHandler configures the mux to serve the "storage" service
// "download" endpoint.
func MountDownloadHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleStorageOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/storage/package/{aip_id}/download", f)
}

// NewDownloadHandler creates a HTTP handler which loads the HTTP request and
// calls the "storage" service "download" endpoint.
func NewDownloadHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDownloadRequest(mux, decoder)
		encodeResponse = EncodeDownloadResponse(encoder)
		encodeError    = EncodeDownloadError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "download")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountLocationsHandler configures the mux to serve the "storage" service
// "locations" endpoint.
func MountLocationsHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleStorageOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/storage/location", f)
}

// NewLocationsHandler creates a HTTP handler which loads the HTTP request and
// calls the "storage" service "locations" endpoint.
func NewLocationsHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeLocationsResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "locations")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountAddLocationHandler configures the mux to serve the "storage" service
// "add_location" endpoint.
func MountAddLocationHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleStorageOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/storage/location", f)
}

// NewAddLocationHandler creates a HTTP handler which loads the HTTP request
// and calls the "storage" service "add_location" endpoint.
func NewAddLocationHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeAddLocationRequest(mux, decoder)
		encodeResponse = EncodeAddLocationResponse(encoder)
		encodeError    = EncodeAddLocationError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "add_location")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountMoveHandler configures the mux to serve the "storage" service "move"
// endpoint.
func MountMoveHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleStorageOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/storage/package/{aip_id}/store", f)
}

// NewMoveHandler creates a HTTP handler which loads the HTTP request and calls
// the "storage" service "move" endpoint.
func NewMoveHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeMoveRequest(mux, decoder)
		encodeResponse = EncodeMoveResponse(encoder)
		encodeError    = EncodeMoveError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "move")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountMoveStatusHandler configures the mux to serve the "storage" service
// "move_status" endpoint.
func MountMoveStatusHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleStorageOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/storage/package/{aip_id}/store", f)
}

// NewMoveStatusHandler creates a HTTP handler which loads the HTTP request and
// calls the "storage" service "move_status" endpoint.
func NewMoveStatusHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeMoveStatusRequest(mux, decoder)
		encodeResponse = EncodeMoveStatusResponse(encoder)
		encodeError    = EncodeMoveStatusError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "move_status")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountRejectHandler configures the mux to serve the "storage" service
// "reject" endpoint.
func MountRejectHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleStorageOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/storage/package/{aip_id}/reject", f)
}

// NewRejectHandler creates a HTTP handler which loads the HTTP request and
// calls the "storage" service "reject" endpoint.
func NewRejectHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeRejectRequest(mux, decoder)
		encodeResponse = EncodeRejectResponse(encoder)
		encodeError    = EncodeRejectError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "reject")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountShowHandler configures the mux to serve the "storage" service "show"
// endpoint.
func MountShowHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleStorageOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/storage/package/{aip_id}", f)
}

// NewShowHandler creates a HTTP handler which loads the HTTP request and calls
// the "storage" service "show" endpoint.
func NewShowHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeShowRequest(mux, decoder)
		encodeResponse = EncodeShowResponse(encoder)
		encodeError    = EncodeShowError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "show")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountShowLocationHandler configures the mux to serve the "storage" service
// "show-location" endpoint.
func MountShowLocationHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleStorageOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/storage/location/{uuid}", f)
}

// NewShowLocationHandler creates a HTTP handler which loads the HTTP request
// and calls the "storage" service "show-location" endpoint.
func NewShowLocationHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeShowLocationRequest(mux, decoder)
		encodeResponse = EncodeShowLocationResponse(encoder)
		encodeError    = EncodeShowLocationError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "show-location")
		ctx = context.WithValue(ctx, goa.ServiceKey, "storage")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service storage.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleStorageOrigin(h)
	mux.Handle("OPTIONS", "/storage/package/{aip_id}/submit", h.ServeHTTP)
	mux.Handle("OPTIONS", "/storage/package/{aip_id}/update", h.ServeHTTP)
	mux.Handle("OPTIONS", "/storage/package/{aip_id}/download", h.ServeHTTP)
	mux.Handle("OPTIONS", "/storage/location", h.ServeHTTP)
	mux.Handle("OPTIONS", "/storage/package/{aip_id}/store", h.ServeHTTP)
	mux.Handle("OPTIONS", "/storage/package/{aip_id}/reject", h.ServeHTTP)
	mux.Handle("OPTIONS", "/storage/package/{aip_id}", h.ServeHTTP)
	mux.Handle("OPTIONS", "/storage/location/{uuid}", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandleStorageOrigin applies the CORS response headers corresponding to the
// origin for the service storage.
func HandleStorageOrigin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOrigin(origin, "*") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS")
			}
			h.ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
		return
	})
}
