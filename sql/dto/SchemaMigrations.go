package dto

type SchemaMigrations struct {
	Version int `db:"version"`
	Dirty   int `db:"dirty"`
}
