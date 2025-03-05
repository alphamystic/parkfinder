package domain

import (
  "database/sql"
)


type Domain struct {
  Dbs *sql.DB
}
