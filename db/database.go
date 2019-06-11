package db

// DataType can be INT, STRING, BOOL, NULL
type DataType int

// Schema represents a collection of DataTypes
// that represent a column within a Table
type Schema map[string]DataType

// Row is an arbitrary tuple of data
type Row *interface{}

// DataType constants
const (
	NULL DataType = iota
	INT
	BOOL
	STRING
)

// Database represents the global wrapper
// for a buoy instance
type Database struct {
	name   string
	tables *[]Table
}

// Exec will execute an SQL Query against a Database
func (db *Database) Exec(query string) ([]Row, error) {
	return nil, nil
}

// Close with gracefully tie up loose ends
func (db *Database) Close() {
}

// Table represents a single entity within
// a Database
type Table struct {
	name   string
	schema *Schema
	rows   *[]Row
}
