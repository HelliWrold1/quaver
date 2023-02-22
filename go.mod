module github.com/HelliWrold1/quaver

go 1.16

require (
	github.com/apache/thrift v0.18.0
	github.com/cloudwego/hertz v0.4.2
	github.com/cloudwego/kitex v0.4.4
	github.com/fsnotify/fsnotify v1.6.0
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/hertz-contrib/jwt v1.0.2
	github.com/hertz-contrib/obs-opentelemetry/logging/logrus v0.1.1
	github.com/hertz-contrib/obs-opentelemetry/tracing v0.1.1
	github.com/hertz-contrib/pprof v0.1.0
	github.com/kitex-contrib/obs-opentelemetry v0.1.0
	github.com/kitex-contrib/obs-opentelemetry/logging/logrus v0.0.0-20230219092456-5e6c84962323
	github.com/kitex-contrib/registry-etcd v0.1.0
	github.com/spf13/viper v1.15.0
	github.com/stretchr/testify v1.8.1
	gorm.io/driver/mysql v1.4.7
	gorm.io/driver/sqlite v1.4.4
	gorm.io/gen v0.3.21
	gorm.io/gorm v1.24.5
	gorm.io/plugin/dbresolver v1.4.1
	gorm.io/plugin/opentelemetry v0.1.0
)

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0
