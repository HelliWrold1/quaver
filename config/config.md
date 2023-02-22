提供获取配置信息的方法，使用前需要初始化，使用viper进行解析

需要将项目目录放到`$GOPATH/src/github.com/HelliWrold1/`下，使[config.go](./config.go)读取的路径为项目目录`$GOPATH/src/github.com/HelliWrold1/quaver`

```shell
cd $GOPATH
mkdir -p src/github.com/HelliWrold1
cd ./src/github.com/HelliWrold1
git clone git@github.com:HelliWrold1/quaver.git
```

### 使用示例
```go
package main

import (
	"github.com/HelliWrold1/quaver/config"
	"fmt"
)
func main() {
	conf := config.NewQuaverConfig()
	conf.LocalConfigInit()
	fmt.Println(conf.MySQLConfig.Host)
}

```