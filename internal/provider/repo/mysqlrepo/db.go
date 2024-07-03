package mysqlrepo

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/danielmesquitta/products-api/internal/config"
	"github.com/danielmesquitta/products-api/internal/provider/db/mysqldb"
)

func NewMySQLDBConn(
	env *config.Env,
) *mysqldb.Queries {
	dbConn, err := sql.Open(
		"mysql",
		env.DBConnection,
	)
	if err != nil {
		panic(err)
	}

	err = dbConn.Ping()
	if err != nil {
		panic(err)
	}

	return mysqldb.New(dbConn)
}
