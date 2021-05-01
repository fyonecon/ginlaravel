package driver
// mysql文档：http://www.topgoer.com/%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C/go%E6%93%8D%E4%BD%9Cmysql/mysql%E4%BD%BF%E7%94%A8.html

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
	"ginlaravel/app/Kit"
	"ginlaravel/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"time"
)

var MysqlDb *sql.DB
var MysqlDbErr error

func init() {
	log.Println("尝试连接MySQL服务...")

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

	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		log.Println("MySQL数据库连接失败。。。")
		log.Println(MysqlDbErr.Error())
		Kit.Log(">>> MySQL数据库连接失败。", "")
		//os.Exit(200)
	}else {
		log.Println("MySQL已连接 >>> ")
	}
}
