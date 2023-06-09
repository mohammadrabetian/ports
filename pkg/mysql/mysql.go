package mysql

import (
	"database/sql"
	"fmt"

	"github.com/mohammadrabetian/ports/util"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SQLDatabase struct {
	*gorm.DB
	SQLDB *sql.DB
}

func NewDatabase(config util.Config) SQLDatabase {

	username := config.MySQL.User
	password := config.MySQL.Password
	host := config.MySQL.Host
	port := config.MySQL.Port
	dbname := config.MySQL.DBName

	url := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		panic("failed to connect to mysql database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to obtain *sql.DB from GORM")
	}

	logrus.Info("SQLDatabase connection established")

	return SQLDatabase{
		DB:    db,
		SQLDB: sqlDB,
	}
}

func (db SQLDatabase) Close() error {
	return db.SQLDB.Close()
}
