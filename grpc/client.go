package grpc

import (
	"google.golang.org/grpc"
)

func Client(toService string) (*grpc.ClientConn, error) {
	//conn, err := consul.AgentHealthService(toService)
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println(conn)
	return grpc.Dial("consul://10.2.171.85:8500/"+toService+"?wait=14s", grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"LoadBalancingPolicy": "round_robin"}`))
}
