package cfg

import "database/sql"

type Config struct {
	DB *sql.DB
}

var CFG = Config{}
