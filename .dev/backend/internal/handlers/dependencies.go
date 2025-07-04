package handlers

import (
	"gameReport/internal/core"
	"gameReport/internal/infrastructure/db"
)

var (
	gameRegistry core.GameRegistry
	dbConn *db.DB
)
// SetGameRegistry sets the global game registry
// TODO: Replace with proper dependency injection
func SetDependencies(registry core.GameRegistry, db *db.DB) {
	gameRegistry = registry
	dbConn = db
}