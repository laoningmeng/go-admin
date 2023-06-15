package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func main() {
	//create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("192.168.2.4", 8848),
	}

	//create ClientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "97d5a42d-e09d-4abf-81c8-e18f1fb108e2", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
		Username:            "nacos",
		Password:            "nacos",
	}

	// create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "config-dev.yaml",
		Group:  "dev",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("content:", content)
}
