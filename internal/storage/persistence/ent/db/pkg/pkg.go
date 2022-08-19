// Code generated by ent, DO NOT EDIT.

package pkg

import (
	"fmt"

	"github.com/artefactual-sdps/enduro/internal/storage/status"
)

const (
	// Label holds the string label denoting the pkg type in the database.
	Label = "pkg"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAipID holds the string denoting the aip_id field in the database.
	FieldAipID = "aip_id"
	// FieldLocationID holds the string denoting the location_id field in the database.
	FieldLocationID = "location_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldObjectKey holds the string denoting the object_key field in the database.
	FieldObjectKey = "object_key"
	// EdgeLocation holds the string denoting the location edge name in mutations.
	EdgeLocation = "location"
	// Table holds the table name of the pkg in the database.
	Table = "package"
	// LocationTable is the table that holds the location relation/edge.
	LocationTable = "package"
	// LocationInverseTable is the table name for the Location entity.
	// It exists in this package in order to avoid circular dependency with the "location" package.
	LocationInverseTable = "location"
	// LocationColumn is the table column denoting the location relation/edge.
	LocationColumn = "location_id"
)

// Columns holds all SQL columns for pkg fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldAipID,
	FieldLocationID,
	FieldStatus,
	FieldObjectKey,
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

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s status.PackageStatus) error {
	switch s.String() {
	case "unspecified", "in_review", "rejected", "stored", "moving":
		return nil
	default:
		return fmt.Errorf("pkg: invalid enum value for status field: %q", s)
	}
}