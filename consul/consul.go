package consul

import (
	"fmt"
	capi "github.com/hashicorp/consul/api"
	"strconv"
)

// consul服务发现
func AgentHealthService(serviceName string) (string, error) {
	fmt.Println("我进到consul发现里面了")
	client, err := capi.NewClient(capi.DefaultConfig()) //创建consul客户端
	if err != nil {
		return "", err
	}
	sr, info, err := client.Agent().AgentHealthServiceByName(serviceName) //健康查询
	if err != nil {
		return "", err
	}
	if sr != "passing" { //如果健康状态不是 "passing"，则返回一个错误。
		return "", fmt.Errorf("is not have health service")
	}
	for _, v := range info {
		fmt.Println("consul:", v)
	}
	//如果健康状态为 "passing"，则将其中一个健康实例的地址返回。
	return fmt.Sprintf("%v:%v", info[0].Service.Address, info[0].Service.Port), nil
}

// consul服务注册
func ServiceRegister(address, port, ip, ConsulPort string) error {
	client, err := capi.NewClient(capi.DefaultConfig())
	if err != nil {
		return err
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		return err
	}

	capi.DefaultConfig().Address = fmt.Sprintf("%v,%v", address, ConsulPort)

	return client.Agent().ServiceRegister(&capi.AgentServiceRegistration{
		ID:      "test",
		Name:    "user",
		Tags:    []string{"GRPC"},
		Port:    portInt,
		Address: address,
		Check: &capi.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%v:%v", ip, port),
			Interval:                       "5s",
			DeregisterCriticalServiceAfter: "10s",
		},
	})
}
