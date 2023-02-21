package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func Init() {
	workDir, _ := os.Getwd()      // 获取工作目录
	viper.SetConfigName("config") // 配置文件的文件名
	viper.SetConfigType("yml")    // 配置文件的后缀
	viper.AddConfigPath(workDir + "/config")
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
