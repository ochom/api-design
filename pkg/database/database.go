package database

import "github.com/ochom/api-design/pkg/models"

// GetSchema return a list of models that are used to create the SQL database schema
// The models are defined in pkg/models
func GetSchema() []interface{} {
	return []interface{}{
		&models.Ping{},
	}
}
