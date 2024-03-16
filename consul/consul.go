package consul

import (
	"fmt"
	"github.com/Minggyyds/framework/grpc"
	"github.com/google/uuid"
	capi "github.com/hashicorp/consul/api"
	"strconv"
)

// consul服务发现
func AgentHealthService(dataid, serviceName string) (string, error) {
	//return grpc.Dial("consul://10.2.171.70:8500/"+serviceName+"?wait=14s", grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"LoadBalancingPolicy": "round_robin"}`))
	fmt.Println("我进到consul发现里面了")
	cof, err := grpc.GetConfig(dataid)
	if err != nil {
		return "", err
	}
	fmt.Println("我进到consul发现里面了")
	//fmt.Println(address, port, ip, ConsulPort)
	clien := capi.DefaultConfig()
	clien.Address = fmt.Sprintf("%v:%v", cof.App.Address, cof.App.ConsulPort)
	client, err := capi.NewClient(clien) //创建consul客户端
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
func ServiceRegister(ip, port, ConsulPort, address string) error {
	fmt.Println(address, port, ip, ConsulPort)
	clien := capi.DefaultConfig()
	clien.Address = fmt.Sprintf("%v:%v", address, ConsulPort)
	client, err := capi.NewClient(clien)
	if err != nil {
		return err
	}

	portInt, err := strconv.Atoi(port)

	if err != nil {
		return err
	}

	return client.Agent().ServiceRegister(&capi.AgentServiceRegistration{
		ID:      uuid.NewString(),
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

//func ServiceRegister(nacosGroup, serviceName string, address string, port string) error {
//	config := capi.DefaultConfig()
//	config.Address = "10.2.171.85:8500"
//	client, _ := capi.NewClient(config)
//	return client.Agent().ServiceRegister(&capi.AgentServiceRegistration{
//		ID:      uuid.NewString(),
//		Name:    "user",
//		Tags:    []string{"GRPC"},
//		Port:    8001,
//		Address: GetIp()[0],
//		Check: &capi.AgentServiceCheck{
//			GRPC:                           fmt.Sprintf("%v:%v", GetIp()[0], "8001"),
//			Interval:                       "5s",
//			DeregisterCriticalServiceAfter: "10s",
//		},
//	})
//}
//
//func GetIp() (ip []string) {
//	addrs, err := net.InterfaceAddrs()
//	if err != nil {
//		return ip
//	}
//	for _, addr := range addrs {
//		ipNet, isVailIpNet := addr.(*net.IPNet)
//		if isVailIpNet && !ipNet.IP.IsLoopback() {
//			if ipNet.IP.To4() != nil {
//				ip = append(ip, ipNet.IP.String())
//			}
//		}
//
//	}
//	return ip
//}
