package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type Registry struct {
	Host string
	Port int
}

// 接口
type RegistryClient interface {
	Register(address string, port int, name, id string, tags []string) error
	DeRegister(serviceId string) error
}

// 类似构造函数
func NewRegistryClient(Host string, port int) RegistryClient {
	return &Registry{
		Host: Host,
		Port: port,
	}
}

func (r *Registry) Register(address string, port int, name, id string, tags []string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}
	check := &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d/health", address, port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "1m",
	}
	registration := api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Address: address,
		Port:    port,
		Tags:    tags,
		Check:   check,
	}
	err = client.Agent().ServiceRegister(&registration)
	if err != nil {
		return err
	}
	return nil
}

// 注销服务
func (r *Registry) DeRegister(serviceId string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}
	err = client.Agent().ServiceDeregister(serviceId)
	if err != nil {
		return err
	}
	return nil
}
