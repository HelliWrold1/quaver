package ftp

import (
	"github.com/HelliWrold1/quaver/config"
	"github.com/jlaffaye/ftp"
	"log"
	"time"
)

var MyFTP *ftp.

func InitFTP() {
	// 加载配置文件
	conf := config.NewQuaverConfig()
	conf.LocalConfigInit()
	//获取到ftp的链接
	var err error
	MyFTP, err = ftp.Connect(连接地址)
	if err != nil {
		log.Printf("获取到FTP链接失败！！！")
	}
	log.Printf("获取到FTP链接成功%v：", MyFTP)
	//登录
	err = MyFTP.Login(登录ftp)
	if err != nil {
		log.Printf("FTP登录失败！！！")
	}
	log.Printf("FTP登录成功！！！")
	//维持长链接
	go keepAlive()
}

func keepAlive() {
	time.Sleep(time.Duration(config.HeartbeatTime) * time.Second)
	MyFTP.Noop()
}
