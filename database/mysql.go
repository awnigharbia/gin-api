// database/mysql.go
package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Db *xorm.Engine

// ConnectDB creates a connection to the MySQL database.
func ConnectDB() (*xorm.Engine, error) {
	var err error
	Db, err = xorm.NewEngine("mysql", "u771241703_seba:12341234aaAA@tcp(sql918.main-hosting.eu:3306)/u771241703_seba")

	if err != nil {
		return nil, err
	}

	return Db, nil
}
