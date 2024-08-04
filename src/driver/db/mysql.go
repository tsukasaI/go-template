package db

import (
	"fmt"
	"my-go-template/src/core/config"

	"gorm.io/gorm"
)

type SQLHandler struct {
	Conn *gorm.DB
}

func NewMysqlSQLHandler() *SQLHandler {
	// conn, err := gorm.Open(mysql.Open(getMysqlDSN()))

	// if err != nil {
	// 	panic(err)
	// }
	// return &SqlHandler{Conn: conn}
	fmt.Println(getMysqlDSN())
	return &SQLHandler{Conn: &gorm.DB{}}
}

func getMysqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MySQLUser(),
		config.MySQLPassword(),
		config.MySQLHost(),
		config.MySQLPort(),
		config.MySQLDatabase(),
	)
}
