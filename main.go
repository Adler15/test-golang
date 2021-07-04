package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"stduentProject/app/student"
	"stduentProject/app/test"
	_ "stduentProject/docs"
	"stduentProject/routers"
	"stduentProject/utils"
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

	url := ginSwagger.URL("http://localhost:9195/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	//在环境变量中获取nacos的ip
	nacosIp := os.Getenv("BCF_NACOS_IP")
	// 创建clientConfig的另一种方式
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(""), //当namespace是public时，此处填空字符串。
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		//constant.WithLogDir("/tmp/nacos/log"),
		//constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithRotateTime("1h"),
		constant.WithMaxAge(3),
		constant.WithLogLevel("debug"),
	)

	// 创建serverConfig的另一种方式
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			nacosIp,
			8082,
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos"),
		),
		//*constant.NewServerConfig(
		//"console2.nacos.io",
		//80,
		//constant.WithScheme("http"),
		//constant.WithContextPath("/nacos")
		//),
	}

	// 创建服务发现客户端的另一种方式 (推荐)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		panic(err)
	}

	// 创建动态配置客户端的另一种方式 (推荐)
	//configClient, err := clients.NewConfigClient(
	//	vo.NacosClientParam{
	//		ClientConfig:  &clientConfig,
	//		ServerConfigs: serverConfigs,
	//	},
	//)

	if ip, err2 := utils.ExternalIP();err2 == nil {
		namingClient.RegisterInstance(vo.RegisterInstanceParam{
			Ip:          ip.String(),
			Port:        9195,
			ServiceName: "golang-demo",
			Weight:      10,
			Enable:      true,
			Healthy:     true,
			Ephemeral:   true,
		})
	}

	//ClusterName=DEFAULT,GroupName=DEFAULT_GROUP
	//namingClient.RegisterInstance(vo.RegisterInstanceParam{
	//	Ip:          "192.168.130.109",
	//	Port:        9195,
	//	ServiceName: "cfl.demo.go",
	//	Weight:      10,
	//	Enable:      true,
	//	Healthy:     true,
	//	Ephemeral:   true,
	//	Metadata:    map[string]string{"idc":"chengdu"},
	//})

	if err := r.Run(":9195"); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
