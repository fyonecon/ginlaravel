package Kit
// 定义全局数据库主语法，这样就可以避免空间命名混乱造成的全局数据库主语法定义位置多的问题。
// 引用如：Kit.DB.Table("gl_user").xxx

import (
	"database/sql"
	"ginvel.com/bootstrap/driver"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var Db *sql.DB = driver.MysqlDb // 连接mysql扩展
var DB *gorm.DB = driver.GDB    // 连接gorm扩展
var RDB *redis.Client = driver.RedisDb // 连接Redis扩展
