package database

import (
	"fmt"
	"wx-server/config"
	"xorm.io/xorm"
)

var db *xorm.Engine

func Engin() *xorm.Engine {
	return db
}

func init() {
	user := config.Cfg().Mysql.User
	pwd := config.Cfg().Mysql.Pwd
	host := config.Cfg().Mysql.Host
	port := config.Cfg().Mysql.Port
	dbName := config.Cfg().Mysql.Db
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local", user, pwd, host, port, dbName)
	_db, err := xorm.NewEngine("mysql", uri)
	if err != nil {
		rslog.Error(err)
	} else {
		db = _db
		db.SetLogger(rlog.GetLogger())
		db.ShowSQL(true)
		db.SetMaxIdleConns(config.Cfg().MustInt("db", "mysql.maxIdleCons", 10))
		db.SetMaxOpenConns(config.Cfg().MustInt("db", "mysql.maxOpenCons", 200))
		db.ShowExecTime(true)

	}

}
