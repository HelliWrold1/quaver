package db

import (
	"github.com/HelliWrold1/quaver/config"
	"github.com/HelliWrold1/quaver/dal/query"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"time"
)

var DB *gorm.DB

func Init() {
	dsn := config.GetDSN()
	maxIdleConns, maxOpenConns, maxLifeTime, maxIdleTime := config.GetConnPoolConfig()

	// 设置日志
	gormlogrus := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Millisecond,
			Colorful:      false,
			LogLevel:      logger.Info,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt: true,
			Logger:      gormlogrus,
		},
	)
	if err != nil {
		panic(err)
	}
	query.SetDefault(db)
	// 设置连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTime))
	sqlDB.SetConnMaxIdleTime(time.Duration(maxIdleTime))

	DB = db
	migration()
}
