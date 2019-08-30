package schema

import (
	gql "github.com/mattdamon108/gqlmerge/lib"
)

// NewSchema : returns new schema
func NewSchema() *string {
	schema := gql.Merge("  ", "./schema")

	return schema
}
