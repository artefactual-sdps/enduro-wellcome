// Code generated by goa v3.7.7, DO NOT EDIT.
//
// storage client
//
// Command:
// $ goa-v3.7.7 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package storage

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "storage" service client.
type Client struct {
	SubmitEndpoint     goa.Endpoint
	UpdateEndpoint     goa.Endpoint
	DownloadEndpoint   goa.Endpoint
	ListEndpoint       goa.Endpoint
	MoveEndpoint       goa.Endpoint
	MoveStatusEndpoint goa.Endpoint
	RejectEndpoint     goa.Endpoint
	ShowEndpoint       goa.Endpoint
}

// NewClient initializes a "storage" service client given the endpoints.
func NewClient(submit, update, download, list, move, moveStatus, reject, show goa.Endpoint) *Client {
	return &Client{
		SubmitEndpoint:     submit,
		UpdateEndpoint:     update,
		DownloadEndpoint:   download,
		ListEndpoint:       list,
		MoveEndpoint:       move,
		MoveStatusEndpoint: moveStatus,
		RejectEndpoint:     reject,
		ShowEndpoint:       show,
	}
}

// Submit calls the "submit" endpoint of the "storage" service.
// Submit may return the following errors:
//	- "not_found" (type *StoragePackageNotfound): Storage package not found
//	- "not_available" (type *goa.ServiceError)
//	- "not_valid" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Submit(ctx context.Context, p *SubmitPayload) (res *SubmitResult, err error) {
	var ires interface{}
	ires, err = c.SubmitEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*SubmitResult), nil
}

// Update calls the "update" endpoint of the "storage" service.
// Update may return the following errors:
//	- "not_found" (type *StoragePackageNotfound): Storage package not found
//	- "not_available" (type *goa.ServiceError)
//	- "not_valid" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Update(ctx context.Context, p *UpdatePayload) (err error) {
	_, err = c.UpdateEndpoint(ctx, p)
	return
}

// Download calls the "download" endpoint of the "storage" service.
// Download may return the following errors:
//	- "not_found" (type *StoragePackageNotfound): Storage package not found
//	- error: internal error
func (c *Client) Download(ctx context.Context, p *DownloadPayload) (res []byte, err error) {
	var ires interface{}
	ires, err = c.DownloadEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]byte), nil
}

// List calls the "list" endpoint of the "storage" service.
func (c *Client) List(ctx context.Context) (res StoredLocationCollection, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(StoredLocationCollection), nil
}

// Move calls the "move" endpoint of the "storage" service.
// Move may return the following errors:
//	- "not_found" (type *StoragePackageNotfound): Storage package not found
//	- "not_available" (type *goa.ServiceError)
//	- "not_valid" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Move(ctx context.Context, p *MovePayload) (err error) {
	_, err = c.MoveEndpoint(ctx, p)
	return
}

// MoveStatus calls the "move_status" endpoint of the "storage" service.
// MoveStatus may return the following errors:
//	- "not_found" (type *StoragePackageNotfound): Storage package not found
//	- "failed_dependency" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) MoveStatus(ctx context.Context, p *MoveStatusPayload) (res *MoveStatusResult, err error) {
	var ires interface{}
	ires, err = c.MoveStatusEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*MoveStatusResult), nil
}

// Reject calls the "reject" endpoint of the "storage" service.
// Reject may return the following errors:
//	- "not_found" (type *StoragePackageNotfound): Storage package not found
//	- "not_available" (type *goa.ServiceError)
//	- "not_valid" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Reject(ctx context.Context, p *RejectPayload) (err error) {
	_, err = c.RejectEndpoint(ctx, p)
	return
}

// Show calls the "show" endpoint of the "storage" service.
// Show may return the following errors:
//	- "not_found" (type *StoragePackageNotfound): Storage package not found
//	- error: internal error
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *StoredStoragePackage, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StoredStoragePackage), nil
}
