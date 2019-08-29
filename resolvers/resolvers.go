package resolvers

import (
	"github.com/rakin92/go-gql-starter/db"
)

// Resolvers including query and mutation
type Resolvers struct {
	*db.DB
}
