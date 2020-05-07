package users_db

import (
	"database/sql"
	"fmt"
	"github.com/Nicrii/Project/users-api/configuration"
	_ "github.com/lib/pq"
)

var (
	Db *sql.DB
)

func init() {
	var err error
	config := configuration.Configuration{}
	config.ReadConfiguration()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s",
		config.Postgres.Host, config.Postgres.Port, config.Postgres.User, config.Postgres.Password, config.Postgres.Dbname)

	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = Db.Ping()
	if err != nil {
		panic(err)
	}
}
