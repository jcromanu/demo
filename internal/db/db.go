package db

import (
	"database/sql"
)

type Database struct {
	conn sql.DB
}
