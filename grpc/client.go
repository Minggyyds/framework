package grpc

import (
	"fmt"
	"github.com/Minggyyds/framework/consul"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Client(dataid, toService string) (*grpc.ClientConn, error) {
	//cof, err := GetConfig(toService)
	//if err != nil {
	//	return nil, err
	//}
	conn, err := consul.AgentHealthService(dataid, toService)
	if err != nil {
		return nil, err
	}
	fmt.Println(conn)
	return grpc.Dial(conn, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//return grpc.Dial("10.2.171.85:8500", grpc.WithInsecure(), grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "round_robin", "service": {"method": "%s"}}`, toService)))
	//consulURL := fmt.Sprintf("consul://%s:%s/+%s+?wait=14s", cof.App.Ip, cof.App.ConsulPort, toService)
	//return grpc.Dial(consulURL, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"LoadBalancingPolicy": "round_robin"}`))
}
