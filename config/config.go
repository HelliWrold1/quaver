package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strings"
)

type QuaverConfig struct {
	*ServerConfig  `mapstructure:"server"`
	*MySQLConfig   `mapstructure:"mysql"`
	*EtcdConfig    `mapstructure:"etcd"`
	*RedisConfig   `mapstructure:"redis"`
	*GormGenConfig `mapstructure:"gormgen"`
}

type ServerConfig struct {
	JWTSecret          string `mapstructure:"jwtSecret"`
	IdentityKey        string `mapstructure:"identityKey"`
	SHA256key          string `mapstructure:"sha256key"`
	UserServiceName    string `mapstructure:"userServiceName"`
	UserServiceAddr    string `mapstructure:"userServiceAddr"`
	CommentServiceName string `mapstructure:"CommentServiceName"`
	CommentServiceAddr string `mapstructure:"CommentServiceAddr"`
	ApiServiceName     string `mapstructure:"apiServiceName"`
	ApiServiceAddr     string `mapstructure:"apiServiceAddr"`
	ExportEndpoint     string `mapstructure:"exportEndpoint"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Charset      string `mapstructure:"charset"`
	Database     string `mapstructure:"database"`
	Port         string `mapstructure:"port"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
	MaxLifeTime  int    `mapstructure:"maxLifeTime"`
	MaxIdleTime  int    `mapstructure:"maxIdleTime"`
}

type EtcdConfig struct {
	Address string `mapstructure:"address"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	Password     string `mapstructure:"password"`
	Database     string `mapstructure:"database"`
	PoolSize     int    `mapstructure:"poolSize"`
	MinIdleConns int    `mapstructure:"minIdleConns"`
}

type GormGenConfig struct {
	QueryOutPath string `mapstructure:"queryOutPath"`
}

func NewQuaverConfig() *QuaverConfig {
	return &QuaverConfig{}
}

func (config *QuaverConfig) LocalConfigInit() {
	viper.AutomaticEnv()
	//goPath := viper.GetString("GOPATH")                                // 获取环境变量
	configPath := "/Users/fengdacrcy/Desktop/quater_chen/quaver/config" // configPath是配置文件的绝对路径
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(config)
	if err != nil {
		fmt.Println(err.Error())
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config changed")
		err := viper.Unmarshal(config)
		if err != nil {
			fmt.Println(err.Error())
		}
	})
}

func (config *QuaverConfig) GetMySqlDSN() string {
	host := config.MySQLConfig.Host
	port := config.MySQLConfig.Port
	database := config.MySQLConfig.Database
	username := config.MySQLConfig.Username
	password := config.MySQLConfig.Password
	charset := config.MySQLConfig.Charset
	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/",
		database, "?charset=", charset + "&parseTime=True&loc=Local"}, "")
	return dsn
}

func (config *QuaverConfig) GetMysqlConnPoolConfig() (int, int, int, int) {
	maxIdleConns := config.MySQLConfig.MaxIdleConns
	maxOpenConns := config.MySQLConfig.MaxOpenConns
	maxLifeTime := config.MySQLConfig.MaxLifeTime
	maxIdleTime := config.MySQLConfig.MaxIdleTime
	return maxIdleConns, maxOpenConns, maxLifeTime, maxIdleTime
}

func (config *QuaverConfig) RemoteConfigInit() {
	err := viper.AddRemoteProvider("etcd", "", "/config/config.yml")
	if err != nil {
		fmt.Println(err.Error())
	}
	viper.SetConfigType("yml")
	err = viper.ReadRemoteConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = viper.Unmarshal(config)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = viper.WatchRemoteConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("remote config changed")
		err := viper.Unmarshal(config)
		if err != nil {
			fmt.Println(err.Error())
		}
	})
}
