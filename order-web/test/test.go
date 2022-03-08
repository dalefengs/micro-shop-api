package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func Register(id, name, address string, port int) {
	// 默认配置
	cfg := api.DefaultConfig()
	cfg.Address = "172.17.0.1:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	// 生成服务注册对象
	registation := &api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Address: address,
		Port:    port,
		Check: &api.AgentServiceCheck{
			// 每个服务都要提供一个GET接口返回{code:200,success:true}
			HTTP:                           fmt.Sprintf("http://%s:%d/health", address, port),
			Interval:                       "5s", // 定时检查
			Timeout:                        "5s", // 超时时间
			DeregisterCriticalServiceAfter: "1m", // 服务失效多少秒后注销
		},
	}
	err = client.Agent().ServiceRegister(registation)
	if err != nil {
		panic(err)
	}
	fmt.Println("注册成功")
}

func Deregister(id string) {
	// 默认配置
	cfg := api.DefaultConfig()
	cfg.Address = "172.17.0.1:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	err = client.Agent().ServiceDeregister(id)
	if err != nil {
		panic(err)
	}
	fmt.Println("注销服务成功")
}

// 获取所有服务
func AllServices() {
	// 默认配置
	cfg := api.DefaultConfig()
	cfg.Address = "172.17.0.1:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	services, err := client.Agent().Services()
	if err != nil {
		return
	}
	fmt.Println(services)
}

// 过滤服务名称
func FilterService(name string) {
	// 默认配置
	cfg := api.DefaultConfig()
	cfg.Address = "172.17.0.1:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	services, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, name))
	if err != nil {
		return
	}
	fmt.Println(services)
}

func main() {
	// 不能使用 127.0.0.1 因为 Consul 已经部署到 docker 不能识别
	//Register("mirco-shop-api", "mirco-shop-api", "192.168.200.110", 8021)

	//Deregister("mirco-shop-api")
	//AllServices()
	FilterService("mirco-shop-api")
}
