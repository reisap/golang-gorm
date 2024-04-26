package dto

type SchemaMigrationsEntity struct {
	Version int `db:"version"`
	Dirty   int `db:"dirty"`
}
