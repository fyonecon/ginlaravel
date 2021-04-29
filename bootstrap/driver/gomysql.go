package driver

/*
第二作者Author：fyonecon
博客Blog：https://blog.csdn.net/weixin_41827162/article/details/115712700
Github：https://github.com/fyonecon/ginlaravel
邮箱Email：2652335796@qq.com，ikydee@yahoo.com
微信WeChat：fy66881159
所在城市City：长沙ChangSha
*/

import (
	"database/sql"
	"fmt"
	"ginlaravel/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"time"
)

// query need rows.Close to release db ins
// exec will release automatic
var MysqlDb *sql.DB  // db pool instance
var MysqlDbErr error // db err instance



func init() {
	log.Println("初始化MySQL")

	// get db config
	dbConfig := config.GetMySQLConfig()

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local&timeout=%s",
		dbConfig["DB_USER"],
		dbConfig["DB_PWD"],
		dbConfig["DB_HOST"],
		dbConfig["DB_PORT"],
		dbConfig["DB_NAME"],
		dbConfig["DB_CHARSET"],
		dbConfig["DB_TIMEOUT"],
	)

	// connect and open db connection
	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)

	if MysqlDbErr != nil {
		panic("database data source name error: " + MysqlDbErr.Error())
	}

	// max open connections
	dbMaxOpenConns, _ := strconv.Atoi(dbConfig["DB_MAX_OPEN_CONNS"])
	MysqlDb.SetMaxOpenConns(dbMaxOpenConns)

	// max idle connections
	dbMaxIdleConns, _ := strconv.Atoi(dbConfig["DB_MAX_IDLE_CONNS"])
	MysqlDb.SetMaxIdleConns(dbMaxIdleConns)

	// max lifetime of connection if <=0 will forever
	dbMaxLifetimeConns, _ := strconv.Atoi(dbConfig["DB_MAX_LIFETIME_CONNS"])
	MysqlDb.SetConnMaxLifetime(time.Duration(dbMaxLifetimeConns))

	// check db connection at once avoid connect failed
	// else error will be reported until db first sql operate
	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		panic("database connect failed: " + MysqlDbErr.Error())
	}
}
