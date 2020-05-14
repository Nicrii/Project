package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/Nicrii/Project/oauth-api/src/repository/config"
	_ "github.com/lib/pq"
)

var (
	Db *sql.DB
)

func init() {
	var err error
	configRepository := config.NewConfigRepository()
	configuration, err := configRepository.Get()
	if err != nil {
		panic(err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s",
		configuration.Postgres.Host, configuration.Postgres.Port, configuration.Postgres.User,
		configuration.Postgres.Password, configuration.Postgres.Dbname)

	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = Db.Ping()
	if err != nil {
		panic(err)
	}
}
