package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {

	server := http.Server{
		Addr:        ":8081",
		Handler:     &myHandle{},
		ReadTimeout: 5 * time.Second,
	}
	mux = make(map[string]func(http.ResponseWriter, *http.Request))

	//这里配置路由，可以添加自己的方法去处理对应的路由
	mux["/"] = Index

	log.Println("已为您启动了服务，您可以打开浏览器访问 http://127.0.0.1:8081 ，我会输入访问日志。")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

type myHandle struct{}

func (*myHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Println("请求url:", r.URL.String())
	log.Println("请求方法:", r.Method)

	//解析，默认不解析的，否者r.Form将拿不到数据
	r.ParseForm()

	log.Println("请求报文:", r)
	log.Println("请求的参数:", r.Form)

	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
	} else {
		fileTer(w, r)
	}

}

// 返回的jsonBean
type BaseJsonBean struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 创建jsonBean
func NewBaseJsonBean(code int, message string, data interface{}) *BaseJsonBean {
	return &BaseJsonBean{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// 文件过滤器
func fileTer(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	//判断是否有.
	if strings.Contains(path, ".") {
		request_type := path[strings.LastIndex(path, "."):]
		switch request_type {
		case ".css":
			w.Header().Set("content-type", "text/css; charset=utf-8")
		case ".js":
			w.Header().Set("content-type", "text/javascript; charset=utf-8")
		default:
		}
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Println("获取系统路径失败：", err)
	}

	fin, err := os.Open(wd + path)
	if err != nil {

		log.Println("读取文件失败：", err)
		//关闭文件句柄
		fin.Close()

		//返回json头
		w.Header().Set("content-type", "text/json; charset=utf-8")

		result := NewBaseJsonBean(404, "", "")
		bytes, _ := json.Marshal(result)
		w.Write([]byte(string(bytes)))

		log.Println("返回数据：", string(bytes))

		return
	}

	fd, _ := ioutil.ReadAll(fin)
	w.Write([]byte(fd))

}

// 默认访问方法
func Index(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println("未找到index.html文件，我为您展示默认的首页 :) 快去创建自己的首页吧")

		w.Header().Set("content-type", "text/html; charset=utf-8")
		w.Write([]byte(indeTpl))

	} else {
		t.Execute(w, nil)
	}

}

var indeTpl = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Tiktok</title>
</head>
<body>
	<video src="video/1.mp4" type="video/mp4" controls autoplay width="500px" height="380px" >
	
	</video>
</body>
</html>

`
