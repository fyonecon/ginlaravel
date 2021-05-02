package driver

/*
第二作者Author：fyonecon
博客Blog：https://blog.csdn.net/weixin_41827162/article/details/115712700
Github：https://github.com/fyonecon/ginlaravel
邮箱Email：2652335796@qq.com，ikydee@yahoo.com
微信WeChat：fy66881159
所在城市City：长沙ChangSha
*/

// GORM文档：https://learnku.com/docs/gorm/v2/advanced_query/9757
// GORM文档：https://www.cnblogs.com/zisefeizhu/category/1747066.html

import (
	"database/sql"
	"fmt"
	"ginlaravel/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var GDB *gorm.DB
var gErr error

func init() {

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

	// 连接现有MySQL
	sqlDB, sErr := sql.Open("mysql", dbDSN)
	GDB, gErr = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if sErr != nil {
		log.Println("GORM现有数据库连接失败，GORM功能将不可用。")
	}else {
		log.Println("GORM现有数据库连接成功")
	}

	if gErr != nil {
		log.Println("GORM数据库连接失败")
	}else {
		log.Println("GORM数据库连接成功")
	}


}