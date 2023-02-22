module github.com/HelliWrold1/quaver

go 1.16

require (
	github.com/apache/thrift v0.18.0
	github.com/bytedance/gopkg v0.0.0-20220623074550-9d6d3df70991 // indirect
	github.com/cloudwego/kitex v0.4.4
	github.com/fsnotify/fsnotify v1.6.0
	github.com/kitex-contrib/obs-opentelemetry v0.1.0
	github.com/kitex-contrib/obs-opentelemetry/logging/logrus v0.0.0-20230219092456-5e6c84962323
	github.com/kitex-contrib/registry-etcd v0.1.0
	github.com/spf13/viper v1.15.0
	github.com/stretchr/testify v1.8.1
	github.com/tidwall/gjson v1.14.3 // indirect
	gorm.io/driver/mysql v1.4.7
	gorm.io/driver/sqlite v1.4.4
	gorm.io/gen v0.3.21
	gorm.io/gorm v1.24.5
	gorm.io/plugin/dbresolver v1.4.1
	gorm.io/plugin/opentelemetry v0.1.0
)

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0
