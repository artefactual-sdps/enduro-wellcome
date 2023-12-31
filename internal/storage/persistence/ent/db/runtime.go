// Code generated by ent, DO NOT EDIT.

package db

import (
	"time"

	"github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/location"
	"github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/pkg"
	"github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	locationFields := schema.Location{}.Fields()
	_ = locationFields
	// locationDescCreatedAt is the schema descriptor for created_at field.
	locationDescCreatedAt := locationFields[6].Descriptor()
	// location.DefaultCreatedAt holds the default value on creation for the created_at field.
	location.DefaultCreatedAt = locationDescCreatedAt.Default.(func() time.Time)
	pkgFields := schema.Pkg{}.Fields()
	_ = pkgFields
	// pkgDescCreatedAt is the schema descriptor for created_at field.
	pkgDescCreatedAt := pkgFields[5].Descriptor()
	// pkg.DefaultCreatedAt holds the default value on creation for the created_at field.
	pkg.DefaultCreatedAt = pkgDescCreatedAt.Default.(func() time.Time)
}
