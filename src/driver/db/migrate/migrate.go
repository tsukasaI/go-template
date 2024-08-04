package migrate

import (
	"database/sql"
	"fmt"
	"my-go-template/src/core/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	// migrate
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Execute() {
	db, _ := sql.Open("mysql", getMigrateDSN())
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://src/driver/db/migrate/sql",
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err == migrate.ErrNoChange {
		return
	} else if err != nil {
		panic(err)
	}
}

func getMigrateDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?multiStatements=true",
		config.MySQLUser(),
		config.MySQLPassword(),
		config.MySQLHost(),
		config.MySQLPort(),
		config.MySQLDatabase(),
	)
}
