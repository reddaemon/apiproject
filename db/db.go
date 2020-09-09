package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/reddaemon/apiproject/config"
)

// Get Db connection
func GetDb(c *config.Config) (*sqlx.DB, error) {
	psqlDSN := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Db["host"], c.Db["port"], c.Db["user"], c.Db["pass"], c.Db["name"])

	return sqlx.Connect("postgres", psqlDSN)

}
