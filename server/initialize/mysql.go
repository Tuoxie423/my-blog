package initialize

import (
	"server/global"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// MySQLInit 初始化数据库连接
func MySQLInit() *gorm.DB {
	// 获取数据库配置
	mysqlCfg := global.Config.Mysql
	db, err := gorm.Open(mysql.Open(mysqlCfg.Dsn()), &gorm.Config{
		// 设置日志模式
		Logger: logger.Default.LogMode(mysqlCfg.LogLevel()),
	})
	if err != nil {
		global.Log.Error("MySQL连接失败:", zap.Error(err))
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(mysqlCfg.MaxIdleConns)
	sqlDb.SetMaxOpenConns(mysqlCfg.MaxOpenConns)

	return db
}
