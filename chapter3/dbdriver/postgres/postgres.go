package postgres

import (
	"database/sql"
	"database/sql/driver"
	"errors"
)

type PostgresDriver struct{}

//implements Driver
func (p PostgresDriver) Open(name string) (driver.Conn, error) {
	return nil, errors.New("Unimplemented")
}

var d *PostgresDriver

func init() {
	d = new(PostgresDriver)
	sql.Register("postgresSQL", d)
}
