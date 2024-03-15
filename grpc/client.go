package grpc

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
)

func Client(toService string) (*grpc.ClientConn, error) {

	cof, err := GetConfig(toService)
	if err != nil {
		return nil, err
	}
	//conn, err := consul.AgentHealthService(toService)
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println(conn)
	//return grpc.Dial("10.2.171.85:8500", grpc.WithInsecure(), grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "round_robin", "service": {"method": "%s"}}`, toService)))
	consulURL := fmt.Sprintf("consul://%s:%s/+%s+?wait=14s", cof.App.Ip, cof.App.ConsulPort, toService)
	return grpc.Dial(consulURL, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"LoadBalancingPolicy": "round_robin"}`))
}
