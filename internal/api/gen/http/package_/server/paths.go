// Code generated by goa v3.7.6, DO NOT EDIT.
//
// HTTP request path constructors for the package service.
//
// Command:
// $ goa-v3.7.6 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package server

import (
	"fmt"
)

// MonitorPackagePath returns the URL path to the package service monitor HTTP endpoint.
func MonitorPackagePath() string {
	return "/package/monitor"
}

// ListPackagePath returns the URL path to the package service list HTTP endpoint.
func ListPackagePath() string {
	return "/package"
}

// ShowPackagePath returns the URL path to the package service show HTTP endpoint.
func ShowPackagePath(id uint) string {
	return fmt.Sprintf("/package/%v", id)
}

// DeletePackagePath returns the URL path to the package service delete HTTP endpoint.
func DeletePackagePath(id uint) string {
	return fmt.Sprintf("/package/%v", id)
}

// CancelPackagePath returns the URL path to the package service cancel HTTP endpoint.
func CancelPackagePath(id uint) string {
	return fmt.Sprintf("/package/%v/cancel", id)
}

// RetryPackagePath returns the URL path to the package service retry HTTP endpoint.
func RetryPackagePath(id uint) string {
	return fmt.Sprintf("/package/%v/retry", id)
}

// WorkflowPackagePath returns the URL path to the package service workflow HTTP endpoint.
func WorkflowPackagePath(id uint) string {
	return fmt.Sprintf("/package/%v/workflow", id)
}

// DownloadPackagePath returns the URL path to the package service download HTTP endpoint.
func DownloadPackagePath(id uint) string {
	return fmt.Sprintf("/package/%v/download", id)
}

// BulkPackagePath returns the URL path to the package service bulk HTTP endpoint.
func BulkPackagePath() string {
	return "/package/bulk"
}

// BulkStatusPackagePath returns the URL path to the package service bulk_status HTTP endpoint.
func BulkStatusPackagePath() string {
	return "/package/bulk"
}

// PreservationActionsPackagePath returns the URL path to the package service preservation-actions HTTP endpoint.
func PreservationActionsPackagePath(id uint) string {
	return fmt.Sprintf("/package/%v/preservation-actions", id)
}