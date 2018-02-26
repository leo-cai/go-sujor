package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sujor.com/leo/sujor-api/config"
)

var SqlDB *sql.DB

func init() {
	// sql open 尚未真正建立连接
	var err error
	SqlDB, err = sql.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	// open异常处理
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			SqlDB.Close()
			return
		}
	}()
	if err != nil {
		panic(config.ConnectingError)
	}
	// ping 尝试建立连接
	err = SqlDB.Ping()
	// 连接异常处理
	if err != nil {
		panic(config.PingError)
	}
	// 设置最大连接数
	SqlDB.SetMaxIdleConns(config.DBConfig.MaxIdleConns)
	SqlDB.SetMaxOpenConns(config.DBConfig.MaxOpenConns)
}
