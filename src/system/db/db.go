package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func Connect() (db *xorm.Engine, err error) {
	return xorm.NewEngine("mysql", "root:root@tcp(localhost:3306)/sql?charset=utf8")
	//db, err = xorm.NewEngine("mysql", "root:root@tcp(localhost:3306)/sql?charset=utf8")
	//if err != nil {
	//	println("Here DB", err)
	//	panic(err)
	//}
	//db.DBMetas()
	//return
}
