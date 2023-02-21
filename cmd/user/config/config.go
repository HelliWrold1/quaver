package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

func Init() {
	viper.SetConfigName("../config/config") // 配置文件的文件名
	viper.SetConfigType("yml")              // 配置文件的后缀
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func GetDSN() string {
	host := viper.GetString("mysql.host")
	fmt.Println("host=", host)
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")
	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/",
		database, "?charset=", charset + "&parseTime=True&loc=Local"}, "")
	return dsn
}

func GetConnPoolConfig() (int, int, int, int) {
	maxIdleConns := viper.GetInt("mysql.maxIdleConns")
	maxOpenConns := viper.GetInt("mysql.maxOpenConns")
	maxLifeTime := viper.GetInt("mysql.maxLifeTime")
	maxIdleTime := viper.GetInt("mysql.maxIdleTime")
	return maxIdleConns, maxOpenConns, maxLifeTime, maxIdleTime
}
