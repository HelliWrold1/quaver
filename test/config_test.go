package test

import (
	"github.com/HelliWrold1/quaver/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	conf := config.NewQuaverConfig()
	conf.LocalConfigInit()
	t.Run("MySQLConfig", func(t *testing.T) {
		assert.Equal(t, "127.0.0.1", conf.MySQLConfig.Host)
	})

	t.Run("RedisConfig", func(t *testing.T) {
		assert.Equal(t, "127.0.0.1:6379", conf.RedisConfig.Host)
	})

	t.Run("EtcdConfig", func(t *testing.T) {
		assert.Equal(t, "127.0.0.1:2379", conf.EtcdConfig.Address)
	})

	t.Run("GormGenConfig", func(t *testing.T) {
		assert.Equal(t, "dal/query", conf.GormGenConfig.QueryOutPath)
	})

	t.Run("GetMysqlDSN", func(t *testing.T) {
		assert.Equal(t,
			"root:root@tcp(127.0.0.1:3306)/quaver?charset=utf8mb4&parseTime=True&loc=Local",
			conf.GetMySqlDSN())
	})

	t.Run("GetMysqlConnPoolConfig", func(t *testing.T) {
		maxIdleConns, maxOpenConns, maxLifeTime, maxIdleTime := conf.GetMysqlConnPoolConfig()
		assert.Equal(t, 28, maxIdleConns)
		assert.Equal(t, 100, maxOpenConns)
		assert.Equal(t, 30, maxLifeTime)
		assert.Equal(t, 30, maxIdleTime)
	})
}
