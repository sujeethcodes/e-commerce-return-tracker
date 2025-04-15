package repository

import (
	"fmt"
	"os"

	"github.com/matryer/resync"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var onceMysql resync.Once
var mysqlConn *gorm.DB

type MysqlCon struct {
	Connection *gorm.DB
}

// Singleton pattern
func SingleMysqlConnection() *MysqlCon {
	onceMysql.Do(func() {
		fmt.Println("Enetr DB connection")
		userName := os.Getenv("MYSQL_USERNAME")
		password := os.Getenv("MYSQL_PASSWORD")
		dbHost := os.Getenv("MYSQL_HOST")
		dbName := os.Getenv("MYSQL_DB_NAME")
		dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, dbHost, dbName)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Errorf("could not create db")
		}
		mysqlConn = db
		fmt.Println("DB connection created successfuly")
	})
	return &MysqlCon{Connection: mysqlConn}

}
