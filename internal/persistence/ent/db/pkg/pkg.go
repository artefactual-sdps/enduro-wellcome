// Code generated by ent, DO NOT EDIT.

package pkg

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the pkg type in the database.
	Label = "pkg"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldWorkflowID holds the string denoting the workflow_id field in the database.
	FieldWorkflowID = "workflow_id"
	// FieldRunID holds the string denoting the run_id field in the database.
	FieldRunID = "run_id"
	// FieldAipID holds the string denoting the aip_id field in the database.
	FieldAipID = "aip_id"
	// FieldLocationID holds the string denoting the location_id field in the database.
	FieldLocationID = "location_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldStartedAt holds the string denoting the started_at field in the database.
	FieldStartedAt = "started_at"
	// FieldCompletedAt holds the string denoting the completed_at field in the database.
	FieldCompletedAt = "completed_at"
	// EdgePreservationActions holds the string denoting the preservation_actions edge name in mutations.
	EdgePreservationActions = "preservation_actions"
	// Table holds the table name of the pkg in the database.
	Table = "package"
	// PreservationActionsTable is the table that holds the preservation_actions relation/edge.
	PreservationActionsTable = "preservation_action"
	// PreservationActionsInverseTable is the table name for the PreservationAction entity.
	// It exists in this package in order to avoid circular dependency with the "preservationaction" package.
	PreservationActionsInverseTable = "preservation_action"
	// PreservationActionsColumn is the table column denoting the preservation_actions relation/edge.
	PreservationActionsColumn = "package_id"
)

// Columns holds all SQL columns for pkg fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldWorkflowID,
	FieldRunID,
	FieldAipID,
	FieldLocationID,
	FieldStatus,
	FieldCreatedAt,
	FieldStartedAt,
	FieldCompletedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the Pkg queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByWorkflowID orders the results by the workflow_id field.
func ByWorkflowID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWorkflowID, opts...).ToFunc()
}

// ByRunID orders the results by the run_id field.
func ByRunID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRunID, opts...).ToFunc()
}

// ByAipID orders the results by the aip_id field.
func ByAipID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAipID, opts...).ToFunc()
}

// ByLocationID orders the results by the location_id field.
func ByLocationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocationID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByStartedAt orders the results by the started_at field.
func ByStartedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartedAt, opts...).ToFunc()
}

// ByCompletedAt orders the results by the completed_at field.
func ByCompletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCompletedAt, opts...).ToFunc()
}

// ByPreservationActionsCount orders the results by preservation_actions count.
func ByPreservationActionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPreservationActionsStep(), opts...)
	}
}

// ByPreservationActions orders the results by preservation_actions terms.
func ByPreservationActions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPreservationActionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPreservationActionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PreservationActionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PreservationActionsTable, PreservationActionsColumn),
	)
}