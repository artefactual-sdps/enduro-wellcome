package package_

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	temporalsdk_client "go.temporal.io/sdk/client"

	"github.com/artefactual-sdps/enduro/internal/api/auth"
	goapackage "github.com/artefactual-sdps/enduro/internal/api/gen/package_"
	"github.com/artefactual-sdps/enduro/internal/event"
)

type Service interface {
	// Goa returns an implementation of the goapackage Service.
	Goa() goapackage.Service
	Create(context.Context, *Package) error
	UpdateWorkflowStatus(ctx context.Context, ID uint, name string, workflowID, runID, aipID string, status Status, storedAt time.Time) error
	SetStatus(ctx context.Context, ID uint, status Status) error
	SetStatusInProgress(ctx context.Context, ID uint, startedAt time.Time) error
	SetStatusPending(ctx context.Context, ID uint) error
	SetLocationID(ctx context.Context, ID uint, locationID uuid.UUID) error
	CreatePreservationAction(ctx context.Context, pa *PreservationAction) error
	SetPreservationActionStatus(ctx context.Context, ID uint, status PreservationActionStatus) error
	CompletePreservationAction(ctx context.Context, ID uint, status PreservationActionStatus, completedAt time.Time) error
	CreatePreservationTask(ctx context.Context, pt *PreservationTask) error
	CompletePreservationTask(ctx context.Context, ID uint, status PreservationTaskStatus, completedAt time.Time, note *string) error
}

type packageImpl struct {
	logger         logr.Logger
	db             *sqlx.DB
	tc             temporalsdk_client.Client
	evsvc          event.EventService
	tokenVerifier  auth.TokenVerifier
	ticketProvider *auth.TicketProvider
}

var _ Service = (*packageImpl)(nil)

func NewService(logger logr.Logger, db *sql.DB, tc temporalsdk_client.Client, evsvc event.EventService, tokenVerifier auth.TokenVerifier, ticketProvider *auth.TicketProvider) *packageImpl {
	return &packageImpl{
		logger:         logger,
		db:             sqlx.NewDb(db, "mysql"),
		tc:             tc,
		evsvc:          evsvc,
		tokenVerifier:  tokenVerifier,
		ticketProvider: ticketProvider,
	}
}

func (svc *packageImpl) Goa() goapackage.Service {
	return &goaWrapper{
		packageImpl: svc,
	}
}

func (svc *packageImpl) Create(ctx context.Context, pkg *Package) error {
	query := `INSERT INTO package (name, workflow_id, run_id, aip_id, location_id, status) VALUES (?, ?, ?, ?, ?, ?)`
	args := []interface{}{
		pkg.Name,
		pkg.WorkflowID,
		pkg.RunID,
		pkg.AIPID,
		pkg.LocationID,
		pkg.Status,
	}

	res, err := svc.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error inserting package: %w", err)
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return fmt.Errorf("error retrieving insert ID: %w", err)
	}

	pkg.ID = uint(id)

	if pkg, err := svc.Goa().Show(ctx, &goapackage.ShowPayload{ID: uint(id)}); err == nil {
		ev := &goapackage.EnduroPackageCreatedEvent{ID: uint(id), Item: pkg}
		event.PublishEvent(ctx, svc.evsvc, ev)
	}

	return nil
}

func (svc *packageImpl) UpdateWorkflowStatus(ctx context.Context, ID uint, name string, workflowID, runID, aipID string, status Status, storedAt time.Time) error {
	// Ensure that storedAt is reset during retries.
	completedAt := &storedAt
	if status == StatusInProgress {
		completedAt = nil
	}
	if completedAt != nil && completedAt.IsZero() {
		completedAt = nil
	}

	query := `UPDATE package SET name = ?, workflow_id = ?, run_id = ?, aip_id = ?, status = ?, completed_at = ? WHERE id = ?`
	args := []interface{}{
		name,
		workflowID,
		runID,
		aipID,
		status,
		completedAt,
		ID,
	}

	if err := svc.updateRow(ctx, query, args); err != nil {
		return err
	}

	if pkg, err := svc.Goa().Show(ctx, &goapackage.ShowPayload{ID: ID}); err == nil {
		ev := &goapackage.EnduroPackageUpdatedEvent{ID: uint(ID), Item: pkg}
		event.PublishEvent(ctx, svc.evsvc, ev)
	}

	return nil
}

func (svc *packageImpl) SetStatus(ctx context.Context, ID uint, status Status) error {
	query := `UPDATE package SET status = ? WHERE id = ?`
	args := []interface{}{
		status,
		ID,
	}

	if err := svc.updateRow(ctx, query, args); err != nil {
		return err
	}

	ev := &goapackage.EnduroPackageStatusUpdatedEvent{ID: uint(ID), Status: status.String()}
	event.PublishEvent(ctx, svc.evsvc, ev)

	return nil
}

func (svc *packageImpl) SetStatusInProgress(ctx context.Context, ID uint, startedAt time.Time) error {
	var query string
	args := []interface{}{StatusInProgress}

	if !startedAt.IsZero() {
		query = `UPDATE package SET status = ?, started_at = ? WHERE id = ?`
		args = append(args, startedAt, ID)
	} else {
		query = `UPDATE package SET status = ? WHERE id = ?`
		args = append(args, ID)
	}

	if err := svc.updateRow(ctx, query, args); err != nil {
		return err
	}

	ev := &goapackage.EnduroPackageStatusUpdatedEvent{ID: uint(ID), Status: StatusInProgress.String()}
	event.PublishEvent(ctx, svc.evsvc, ev)

	return nil
}

func (svc *packageImpl) SetStatusPending(ctx context.Context, ID uint) error {
	query := `UPDATE package SET status = ?, WHERE id = ?`
	args := []interface{}{
		StatusPending,
		ID,
	}

	if err := svc.updateRow(ctx, query, args); err != nil {
		return err
	}

	ev := &goapackage.EnduroPackageStatusUpdatedEvent{ID: uint(ID), Status: StatusPending.String()}
	event.PublishEvent(ctx, svc.evsvc, ev)

	return nil
}

func (svc *packageImpl) SetLocationID(ctx context.Context, ID uint, locationID uuid.UUID) error {
	query := `UPDATE package SET location_id = ? WHERE id = ?`
	args := []interface{}{
		locationID,
		ID,
	}

	if err := svc.updateRow(ctx, query, args); err != nil {
		return err
	}

	ev := &goapackage.EnduroPackageLocationUpdatedEvent{ID: uint(ID), LocationID: locationID}
	event.PublishEvent(ctx, svc.evsvc, ev)

	return nil
}

func (svc *packageImpl) updateRow(ctx context.Context, query string, args []interface{}) error {
	_, err := svc.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error updating package: %v", err)
	}

	return nil
}

func (svc *packageImpl) read(ctx context.Context, ID uint) (*Package, error) {
	query := "SELECT id, name, workflow_id, run_id, aip_id, location_id, status, CONVERT_TZ(created_at, @@session.time_zone, '+00:00') AS created_at, CONVERT_TZ(started_at, @@session.time_zone, '+00:00') AS started_at, CONVERT_TZ(completed_at, @@session.time_zone, '+00:00') AS completed_at FROM package WHERE id = ?"
	args := []interface{}{ID}
	c := Package{}

	if err := svc.db.GetContext(ctx, &c, query, args...); err != nil {
		return nil, err
	}

	return &c, nil
}
