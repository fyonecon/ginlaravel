package driver
// GORM文档：https://learnku.com/docs/gorm/v2/advanced_query/9757
// GORM文档：https://www.cnblogs.com/zisefeizhu/category/1747066.html

import (
	"database/sql"
	"fmt"
	"ginvel.com/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var GDB *gorm.DB
var gErr error

func InitGorm() {

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
		log.Println("GORM现有数据库连接失败，GORM功能将不可用。。。", sErr)
		//os.Exit(200)
	}else {
		log.Println("尝试连接GORM... ")
	}

	if gErr != nil {
		log.Println("GORM数据库连接失败。。。", gErr)
		//os.Exit(200)
	}else {
		log.Println("GORM已连接现有数据库驱动 >>> ")
	}

}