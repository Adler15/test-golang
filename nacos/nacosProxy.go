package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"stduentProject/utils"
)

func RegisterNacos(nacosIp string) {

	// 创建clientConfig的另一种方式
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(""), //当namespace是public时，此处填空字符串。
		constant.WithTimeoutMs(5000),
		constant.WithBeatInterval(15000),
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
			ServiceName: "cfl-golang-demo",
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
	//	ServiceName: "cfl-test-go",
	//	Weight:      10,
	//	Enable:      true,
	//	Healthy:     true,
	//	Ephemeral:   true,
	//	Metadata:    map[string]string{"idc": "chengdu"},
	//})

}

