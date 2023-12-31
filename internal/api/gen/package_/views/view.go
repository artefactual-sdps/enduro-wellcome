// Code generated by goa v3.13.2, DO NOT EDIT.
//
// package views
//
// Command:
// $ goa gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package views

import (
	"github.com/google/uuid"
	goa "goa.design/goa/v3/pkg"
)

// EnduroMonitorEvent is the viewed result type that is projected based on a
// view.
type EnduroMonitorEvent struct {
	// Type to project
	Projected *EnduroMonitorEventView
	// View to render
	View string
}

// EnduroStoredPackage is the viewed result type that is projected based on a
// view.
type EnduroStoredPackage struct {
	// Type to project
	Projected *EnduroStoredPackageView
	// View to render
	View string
}

// EnduroPackagePreservationActions is the viewed result type that is projected
// based on a view.
type EnduroPackagePreservationActions struct {
	// Type to project
	Projected *EnduroPackagePreservationActionsView
	// View to render
	View string
}

// EnduroMonitorEventView is a type that runs validations on a projected type.
type EnduroMonitorEventView struct {
	MonitorPingEvent               *EnduroMonitorPingEventView
	PackageCreatedEvent            *EnduroPackageCreatedEventView
	PackageUpdatedEvent            *EnduroPackageUpdatedEventView
	PackageStatusUpdatedEvent      *EnduroPackageStatusUpdatedEventView
	PackageLocationUpdatedEvent    *EnduroPackageLocationUpdatedEventView
	PreservationActionCreatedEvent *EnduroPreservationActionCreatedEventView
	PreservationActionUpdatedEvent *EnduroPreservationActionUpdatedEventView
	PreservationTaskCreatedEvent   *EnduroPreservationTaskCreatedEventView
	PreservationTaskUpdatedEvent   *EnduroPreservationTaskUpdatedEventView
}

// EnduroMonitorPingEventView is a type that runs validations on a projected
// type.
type EnduroMonitorPingEventView struct {
	Message *string
}

// EnduroPackageCreatedEventView is a type that runs validations on a projected
// type.
type EnduroPackageCreatedEventView struct {
	// Identifier of package
	ID   *uint
	Item *EnduroStoredPackageView
}

// EnduroStoredPackageView is a type that runs validations on a projected type.
type EnduroStoredPackageView struct {
	// Identifier of package
	ID *uint
	// Name of the package
	Name *string
	// Identifier of storage location
	LocationID *uuid.UUID
	// Status of the package
	Status *string
	// Identifier of processing workflow
	WorkflowID *string
	// Identifier of latest processing workflow run
	RunID *string
	// Identifier of AIP
	AipID *string
	// Creation datetime
	CreatedAt *string
	// Start datetime
	StartedAt *string
	// Completion datetime
	CompletedAt *string
}

// EnduroPackageUpdatedEventView is a type that runs validations on a projected
// type.
type EnduroPackageUpdatedEventView struct {
	// Identifier of package
	ID   *uint
	Item *EnduroStoredPackageView
}

// EnduroPackageStatusUpdatedEventView is a type that runs validations on a
// projected type.
type EnduroPackageStatusUpdatedEventView struct {
	// Identifier of package
	ID     *uint
	Status *string
}

// EnduroPackageLocationUpdatedEventView is a type that runs validations on a
// projected type.
type EnduroPackageLocationUpdatedEventView struct {
	// Identifier of package
	ID *uint
	// Identifier of storage location
	LocationID *uuid.UUID
}

// EnduroPreservationActionCreatedEventView is a type that runs validations on
// a projected type.
type EnduroPreservationActionCreatedEventView struct {
	// Identifier of preservation action
	ID   *uint
	Item *EnduroPackagePreservationActionView
}

// EnduroPackagePreservationActionView is a type that runs validations on a
// projected type.
type EnduroPackagePreservationActionView struct {
	ID          *uint
	WorkflowID  *string
	Type        *string
	Status      *string
	StartedAt   *string
	CompletedAt *string
	Tasks       EnduroPackagePreservationTaskCollectionView
	PackageID   *uint
}

// EnduroPackagePreservationTaskCollectionView is a type that runs validations
// on a projected type.
type EnduroPackagePreservationTaskCollectionView []*EnduroPackagePreservationTaskView

// EnduroPackagePreservationTaskView is a type that runs validations on a
// projected type.
type EnduroPackagePreservationTaskView struct {
	ID                   *uint
	TaskID               *string
	Name                 *string
	Status               *string
	StartedAt            *string
	CompletedAt          *string
	Note                 *string
	PreservationActionID *uint
}

// EnduroPreservationActionUpdatedEventView is a type that runs validations on
// a projected type.
type EnduroPreservationActionUpdatedEventView struct {
	// Identifier of preservation action
	ID   *uint
	Item *EnduroPackagePreservationActionView
}

// EnduroPreservationTaskCreatedEventView is a type that runs validations on a
// projected type.
type EnduroPreservationTaskCreatedEventView struct {
	// Identifier of preservation task
	ID   *uint
	Item *EnduroPackagePreservationTaskView
}

// EnduroPreservationTaskUpdatedEventView is a type that runs validations on a
// projected type.
type EnduroPreservationTaskUpdatedEventView struct {
	// Identifier of preservation task
	ID   *uint
	Item *EnduroPackagePreservationTaskView
}

// EnduroPackagePreservationActionsView is a type that runs validations on a
// projected type.
type EnduroPackagePreservationActionsView struct {
	Actions EnduroPackagePreservationActionCollectionView
}

// EnduroPackagePreservationActionCollectionView is a type that runs
// validations on a projected type.
type EnduroPackagePreservationActionCollectionView []*EnduroPackagePreservationActionView

var (
	// EnduroMonitorEventMap is a map indexing the attribute names of
	// EnduroMonitorEvent by view name.
	EnduroMonitorEventMap = map[string][]string{
		"default": {
			"monitor_ping_event",
			"package_created_event",
			"package_updated_event",
			"package_status_updated_event",
			"package_location_updated_event",
			"preservation_action_created_event",
			"preservation_action_updated_event",
			"preservation_task_created_event",
			"preservation_task_updated_event",
		},
	}
	// EnduroStoredPackageMap is a map indexing the attribute names of
	// EnduroStoredPackage by view name.
	EnduroStoredPackageMap = map[string][]string{
		"default": {
			"id",
			"name",
			"location_id",
			"status",
			"workflow_id",
			"run_id",
			"aip_id",
			"created_at",
			"started_at",
			"completed_at",
		},
	}
	// EnduroPackagePreservationActionsMap is a map indexing the attribute names of
	// EnduroPackagePreservationActions by view name.
	EnduroPackagePreservationActionsMap = map[string][]string{
		"default": {
			"actions",
		},
	}
	// EnduroMonitorPingEventMap is a map indexing the attribute names of
	// EnduroMonitorPingEvent by view name.
	EnduroMonitorPingEventMap = map[string][]string{
		"default": {
			"message",
		},
	}
	// EnduroPackageCreatedEventMap is a map indexing the attribute names of
	// EnduroPackageCreatedEvent by view name.
	EnduroPackageCreatedEventMap = map[string][]string{
		"default": {
			"id",
			"item",
		},
	}
	// EnduroPackageUpdatedEventMap is a map indexing the attribute names of
	// EnduroPackageUpdatedEvent by view name.
	EnduroPackageUpdatedEventMap = map[string][]string{
		"default": {
			"id",
			"item",
		},
	}
	// EnduroPackageStatusUpdatedEventMap is a map indexing the attribute names of
	// EnduroPackageStatusUpdatedEvent by view name.
	EnduroPackageStatusUpdatedEventMap = map[string][]string{
		"default": {
			"id",
			"status",
		},
	}
	// EnduroPackageLocationUpdatedEventMap is a map indexing the attribute names
	// of EnduroPackageLocationUpdatedEvent by view name.
	EnduroPackageLocationUpdatedEventMap = map[string][]string{
		"default": {
			"id",
			"location_id",
		},
	}
	// EnduroPreservationActionCreatedEventMap is a map indexing the attribute
	// names of EnduroPreservationActionCreatedEvent by view name.
	EnduroPreservationActionCreatedEventMap = map[string][]string{
		"default": {
			"id",
			"item",
		},
	}
	// EnduroPackagePreservationActionMap is a map indexing the attribute names of
	// EnduroPackagePreservationAction by view name.
	EnduroPackagePreservationActionMap = map[string][]string{
		"simple": {
			"id",
			"workflow_id",
			"type",
			"status",
			"started_at",
			"completed_at",
			"package_id",
		},
		"default": {
			"id",
			"workflow_id",
			"type",
			"status",
			"started_at",
			"completed_at",
			"tasks",
			"package_id",
		},
	}
	// EnduroPackagePreservationTaskCollectionMap is a map indexing the attribute
	// names of EnduroPackagePreservationTaskCollection by view name.
	EnduroPackagePreservationTaskCollectionMap = map[string][]string{
		"default": {
			"id",
			"task_id",
			"name",
			"status",
			"started_at",
			"completed_at",
			"note",
			"preservation_action_id",
		},
	}
	// EnduroPackagePreservationTaskMap is a map indexing the attribute names of
	// EnduroPackagePreservationTask by view name.
	EnduroPackagePreservationTaskMap = map[string][]string{
		"default": {
			"id",
			"task_id",
			"name",
			"status",
			"started_at",
			"completed_at",
			"note",
			"preservation_action_id",
		},
	}
	// EnduroPreservationActionUpdatedEventMap is a map indexing the attribute
	// names of EnduroPreservationActionUpdatedEvent by view name.
	EnduroPreservationActionUpdatedEventMap = map[string][]string{
		"default": {
			"id",
			"item",
		},
	}
	// EnduroPreservationTaskCreatedEventMap is a map indexing the attribute names
	// of EnduroPreservationTaskCreatedEvent by view name.
	EnduroPreservationTaskCreatedEventMap = map[string][]string{
		"default": {
			"id",
			"item",
		},
	}
	// EnduroPreservationTaskUpdatedEventMap is a map indexing the attribute names
	// of EnduroPreservationTaskUpdatedEvent by view name.
	EnduroPreservationTaskUpdatedEventMap = map[string][]string{
		"default": {
			"id",
			"item",
		},
	}
	// EnduroPackagePreservationActionCollectionMap is a map indexing the attribute
	// names of EnduroPackagePreservationActionCollection by view name.
	EnduroPackagePreservationActionCollectionMap = map[string][]string{
		"simple": {
			"id",
			"workflow_id",
			"type",
			"status",
			"started_at",
			"completed_at",
			"package_id",
		},
		"default": {
			"id",
			"workflow_id",
			"type",
			"status",
			"started_at",
			"completed_at",
			"tasks",
			"package_id",
		},
	}
)

// ValidateEnduroMonitorEvent runs the validations defined on the viewed result
// type EnduroMonitorEvent.
func ValidateEnduroMonitorEvent(result *EnduroMonitorEvent) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateEnduroMonitorEventView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateEnduroStoredPackage runs the validations defined on the viewed
// result type EnduroStoredPackage.
func ValidateEnduroStoredPackage(result *EnduroStoredPackage) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateEnduroStoredPackageView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateEnduroPackagePreservationActions runs the validations defined on the
// viewed result type EnduroPackagePreservationActions.
func ValidateEnduroPackagePreservationActions(result *EnduroPackagePreservationActions) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateEnduroPackagePreservationActionsView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateEnduroMonitorEventView runs the validations defined on
// EnduroMonitorEventView using the "default" view.
func ValidateEnduroMonitorEventView(result *EnduroMonitorEventView) (err error) {

	if result.MonitorPingEvent != nil {
		if err2 := ValidateEnduroMonitorPingEventView(result.MonitorPingEvent); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.PackageCreatedEvent != nil {
		if err2 := ValidateEnduroPackageCreatedEventView(result.PackageCreatedEvent); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.PackageUpdatedEvent != nil {
		if err2 := ValidateEnduroPackageUpdatedEventView(result.PackageUpdatedEvent); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.PackageStatusUpdatedEvent != nil {
		if err2 := ValidateEnduroPackageStatusUpdatedEventView(result.PackageStatusUpdatedEvent); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.PackageLocationUpdatedEvent != nil {
		if err2 := ValidateEnduroPackageLocationUpdatedEventView(result.PackageLocationUpdatedEvent); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.PreservationActionCreatedEvent != nil {
		if err2 := ValidateEnduroPreservationActionCreatedEventView(result.PreservationActionCreatedEvent); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.PreservationActionUpdatedEvent != nil {
		if err2 := ValidateEnduroPreservationActionUpdatedEventView(result.PreservationActionUpdatedEvent); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.PreservationTaskCreatedEvent != nil {
		if err2 := ValidateEnduroPreservationTaskCreatedEventView(result.PreservationTaskCreatedEvent); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.PreservationTaskUpdatedEvent != nil {
		if err2 := ValidateEnduroPreservationTaskUpdatedEventView(result.PreservationTaskUpdatedEvent); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroMonitorPingEventView runs the validations defined on
// EnduroMonitorPingEventView using the "default" view.
func ValidateEnduroMonitorPingEventView(result *EnduroMonitorPingEventView) (err error) {

	return
}

// ValidateEnduroPackageCreatedEventView runs the validations defined on
// EnduroPackageCreatedEventView using the "default" view.
func ValidateEnduroPackageCreatedEventView(result *EnduroPackageCreatedEventView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Item != nil {
		if err2 := ValidateEnduroStoredPackageView(result.Item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroStoredPackageView runs the validations defined on
// EnduroStoredPackageView using the "default" view.
func ValidateEnduroStoredPackageView(result *EnduroStoredPackageView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "result"))
	}
	if result.Status != nil {
		if !(*result.Status == "new" || *result.Status == "in progress" || *result.Status == "done" || *result.Status == "error" || *result.Status == "unknown" || *result.Status == "queued" || *result.Status == "pending" || *result.Status == "abandoned") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []any{"new", "in progress", "done", "error", "unknown", "queued", "pending", "abandoned"}))
		}
	}
	if result.WorkflowID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.workflow_id", *result.WorkflowID, goa.FormatUUID))
	}
	if result.RunID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.run_id", *result.RunID, goa.FormatUUID))
	}
	if result.AipID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.aip_id", *result.AipID, goa.FormatUUID))
	}
	if result.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.created_at", *result.CreatedAt, goa.FormatDateTime))
	}
	if result.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.started_at", *result.StartedAt, goa.FormatDateTime))
	}
	if result.CompletedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.completed_at", *result.CompletedAt, goa.FormatDateTime))
	}
	return
}

// ValidateEnduroPackageUpdatedEventView runs the validations defined on
// EnduroPackageUpdatedEventView using the "default" view.
func ValidateEnduroPackageUpdatedEventView(result *EnduroPackageUpdatedEventView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Item != nil {
		if err2 := ValidateEnduroStoredPackageView(result.Item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroPackageStatusUpdatedEventView runs the validations defined on
// EnduroPackageStatusUpdatedEventView using the "default" view.
func ValidateEnduroPackageStatusUpdatedEventView(result *EnduroPackageStatusUpdatedEventView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.Status != nil {
		if !(*result.Status == "new" || *result.Status == "in progress" || *result.Status == "done" || *result.Status == "error" || *result.Status == "unknown" || *result.Status == "queued" || *result.Status == "pending" || *result.Status == "abandoned") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []any{"new", "in progress", "done", "error", "unknown", "queued", "pending", "abandoned"}))
		}
	}
	return
}

// ValidateEnduroPackageLocationUpdatedEventView runs the validations defined
// on EnduroPackageLocationUpdatedEventView using the "default" view.
func ValidateEnduroPackageLocationUpdatedEventView(result *EnduroPackageLocationUpdatedEventView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.LocationID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("location_id", "result"))
	}
	return
}

// ValidateEnduroPreservationActionCreatedEventView runs the validations
// defined on EnduroPreservationActionCreatedEventView using the "default" view.
func ValidateEnduroPreservationActionCreatedEventView(result *EnduroPreservationActionCreatedEventView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Item != nil {
		if err2 := ValidateEnduroPackagePreservationActionViewSimple(result.Item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroPackagePreservationActionViewSimple runs the validations
// defined on EnduroPackagePreservationActionView using the "simple" view.
func ValidateEnduroPackagePreservationActionViewSimple(result *EnduroPackagePreservationActionView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.WorkflowID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("workflow_id", "result"))
	}
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.StartedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("started_at", "result"))
	}
	if result.Type != nil {
		if !(*result.Type == "create-aip" || *result.Type == "create-and-review-aip" || *result.Type == "move-package") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.type", *result.Type, []any{"create-aip", "create-and-review-aip", "move-package"}))
		}
	}
	if result.Status != nil {
		if !(*result.Status == "unspecified" || *result.Status == "in progress" || *result.Status == "done" || *result.Status == "error" || *result.Status == "queued" || *result.Status == "pending") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []any{"unspecified", "in progress", "done", "error", "queued", "pending"}))
		}
	}
	if result.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.started_at", *result.StartedAt, goa.FormatDateTime))
	}
	if result.CompletedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.completed_at", *result.CompletedAt, goa.FormatDateTime))
	}
	return
}

// ValidateEnduroPackagePreservationActionView runs the validations defined on
// EnduroPackagePreservationActionView using the "default" view.
func ValidateEnduroPackagePreservationActionView(result *EnduroPackagePreservationActionView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.WorkflowID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("workflow_id", "result"))
	}
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.StartedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("started_at", "result"))
	}
	if result.Type != nil {
		if !(*result.Type == "create-aip" || *result.Type == "create-and-review-aip" || *result.Type == "move-package") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.type", *result.Type, []any{"create-aip", "create-and-review-aip", "move-package"}))
		}
	}
	if result.Status != nil {
		if !(*result.Status == "unspecified" || *result.Status == "in progress" || *result.Status == "done" || *result.Status == "error" || *result.Status == "queued" || *result.Status == "pending") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []any{"unspecified", "in progress", "done", "error", "queued", "pending"}))
		}
	}
	if result.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.started_at", *result.StartedAt, goa.FormatDateTime))
	}
	if result.CompletedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.completed_at", *result.CompletedAt, goa.FormatDateTime))
	}
	if result.Tasks != nil {
		if err2 := ValidateEnduroPackagePreservationTaskCollectionView(result.Tasks); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroPackagePreservationTaskCollectionView runs the validations
// defined on EnduroPackagePreservationTaskCollectionView using the "default"
// view.
func ValidateEnduroPackagePreservationTaskCollectionView(result EnduroPackagePreservationTaskCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateEnduroPackagePreservationTaskView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroPackagePreservationTaskView runs the validations defined on
// EnduroPackagePreservationTaskView using the "default" view.
func ValidateEnduroPackagePreservationTaskView(result *EnduroPackagePreservationTaskView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.TaskID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("task_id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.StartedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("started_at", "result"))
	}
	if result.Status != nil {
		if !(*result.Status == "unspecified" || *result.Status == "in progress" || *result.Status == "done" || *result.Status == "error" || *result.Status == "queued" || *result.Status == "pending") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []any{"unspecified", "in progress", "done", "error", "queued", "pending"}))
		}
	}
	if result.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.started_at", *result.StartedAt, goa.FormatDateTime))
	}
	if result.CompletedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.completed_at", *result.CompletedAt, goa.FormatDateTime))
	}
	return
}

// ValidateEnduroPreservationActionUpdatedEventView runs the validations
// defined on EnduroPreservationActionUpdatedEventView using the "default" view.
func ValidateEnduroPreservationActionUpdatedEventView(result *EnduroPreservationActionUpdatedEventView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Item != nil {
		if err2 := ValidateEnduroPackagePreservationActionViewSimple(result.Item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroPreservationTaskCreatedEventView runs the validations defined
// on EnduroPreservationTaskCreatedEventView using the "default" view.
func ValidateEnduroPreservationTaskCreatedEventView(result *EnduroPreservationTaskCreatedEventView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Item != nil {
		if err2 := ValidateEnduroPackagePreservationTaskView(result.Item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroPreservationTaskUpdatedEventView runs the validations defined
// on EnduroPreservationTaskUpdatedEventView using the "default" view.
func ValidateEnduroPreservationTaskUpdatedEventView(result *EnduroPreservationTaskUpdatedEventView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Item != nil {
		if err2 := ValidateEnduroPackagePreservationTaskView(result.Item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroPackagePreservationActionsView runs the validations defined on
// EnduroPackagePreservationActionsView using the "default" view.
func ValidateEnduroPackagePreservationActionsView(result *EnduroPackagePreservationActionsView) (err error) {

	if result.Actions != nil {
		if err2 := ValidateEnduroPackagePreservationActionCollectionView(result.Actions); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroPackagePreservationActionCollectionViewSimple runs the
// validations defined on EnduroPackagePreservationActionCollectionView using
// the "simple" view.
func ValidateEnduroPackagePreservationActionCollectionViewSimple(result EnduroPackagePreservationActionCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateEnduroPackagePreservationActionViewSimple(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroPackagePreservationActionCollectionView runs the validations
// defined on EnduroPackagePreservationActionCollectionView using the "default"
// view.
func ValidateEnduroPackagePreservationActionCollectionView(result EnduroPackagePreservationActionCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateEnduroPackagePreservationActionView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
