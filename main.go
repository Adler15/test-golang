package main

import (
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"stduentProject/app/student"
	"stduentProject/app/test"
	_ "stduentProject/docs"
	"stduentProject/nacosProxy"
	"stduentProject/routers"
)

//@title The Golang Project of Victor
//@version 1.0
//@description Victor的golang测试项目
//@termsOfService https://github.com/Adler15

//@contact.name Victor
//@contact.url https://github.com/Adler15
//@contact.email chenfeilong1115@hotmail.com

//@license.name Apache 2.0
//@license.url http://www.apache.org/licenses/LICENSE-2.0.html

//@host localhost:9195
//@BasePath
func main() {
	// 加载多个APP的路由配置
	routers.Include(student.Routers,test.Routers)
	// 初始化路由
	r := routers.Init()

	//在环境变量中获取nacos的ip
	nacosIp := os.Getenv("BCF_NACOS_IP")

	nacosProxy.RegisterNacos(nacosIp)

	url := ginSwagger.URL(nacosIp+":9195/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	if err := r.Run(":9195"); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
