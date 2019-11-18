package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs-ink/rslog"
	"wx-server/config"
	"wx-server/rlog"
	"xorm.io/xorm"
)

var db *xorm.Engine

func Engine() *xorm.Engine {
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
		db.SetLogger(&rlog.RXOrmLog{})
		db.ShowSQL(true)
		db.SetMaxIdleConns(config.Cfg().Mysql.MaxIdleConns)
		db.SetMaxOpenConns(config.Cfg().Mysql.MaxOpenConns)
		db.ShowExecTime(true)
	}

}
