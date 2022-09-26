// Code generated by goa v3.8.5, DO NOT EDIT.
//
// upload service
//
// Command:
// $ goa-v3.8.5 gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package upload

import (
	"context"
	"io"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// The upload service handles file submissions to the SIPs bucket.
type Service interface {
	// Upload implements upload.
	Upload(context.Context, *UploadPayload, io.ReadCloser) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// OAuth2Auth implements the authorization logic for the OAuth2 security scheme.
	OAuth2Auth(ctx context.Context, token string, schema *security.OAuth2Scheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "upload"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"upload"}

// UploadPayload is the payload type of the upload service upload method.
type UploadPayload struct {
	// Content-Type header, must define value for multipart boundary.
	ContentType string
	OauthToken  *string
}

// Invalid token
type Unauthorized string

// Error returns an error description.
func (e Unauthorized) Error() string {
	return "Invalid token"
}

// ErrorName returns "unauthorized".
func (e Unauthorized) ErrorName() string {
	return "unauthorized"
}

// MakeInvalidMediaType builds a goa.ServiceError from an error.
func MakeInvalidMediaType(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "invalid_media_type", false, false, false)
}

// MakeInvalidMultipartRequest builds a goa.ServiceError from an error.
func MakeInvalidMultipartRequest(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "invalid_multipart_request", false, false, false)
}

// MakeInternalError builds a goa.ServiceError from an error.
func MakeInternalError(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "internal_error", false, false, false)
}
