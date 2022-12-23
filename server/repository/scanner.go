package repository

// Scanner interface represent a model that can be scanned into values.
// Used by the internal helpers of each store to scan sql.Row or sql.Rows into models.
type Scanner interface {
	Scan(dest ...interface{}) error
}
