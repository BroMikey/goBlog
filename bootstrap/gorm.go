package bootstrap

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitGorm initializes and returns a GORM DB connection.
// It receives a Config pointer and a logrus.Logger for logging.
// Returns nil if MySQL is not configured.
func InitGorm(conf *Config, log *logrus.Logger) *gorm.DB {
	if conf.Mysql.Host == "" {
		log.Warnln("未配置mysql，取消gorm连接")
		return nil
	}

	dsn := conf.Mysql.DSN()

	// Set GORM logger based on environment
	var mysqlLogger logger.Interface
	if conf.System.Env == "debug" || conf.System.Env == "dev" {
		// Development: show all SQL
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		// Production: only show errors
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Fatalf("mysql连接失败 [%s]: %v", dsn, err)
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取底层sql.DB失败: %v", err)
		return nil
	}

	// Connection pool settings (read from config, with defaults)
	maxIdleConns := conf.Mysql.MaxIdleConn
	if maxIdleConns <= 0 {
		maxIdleConns = 10
	}
	maxOpenConns := conf.Mysql.MaxOpenConn
	if maxOpenConns <= 0 {
		maxOpenConns = 100
	}

	sqlDB.SetMaxIdleConns(maxIdleConns)     // 最大空闲连接数
	sqlDB.SetMaxOpenConns(maxOpenConns)     // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间

	log.Infof("mysql连接成功: %s:%d/%s", conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.DBname)

	return db
}
