package Common

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/Mr-Herod/CloudGamingDemo/Naming/naming"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:10086", "the address to connect to")
)

func RegisterServer(serviceName, ip string, port int32) error {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	c := naming.NewNamingServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RegisterService(ctx, &naming.RegisterServiceReq{
		ServiceName: serviceName,
		IP:          ip,
		Port:        port,
	})
	if err != nil {
		return fmt.Errorf("could not RegisterService: %v", err)
	}
	log.Printf("rsp: %v", r)
	return nil
}

func FindServer(serviceName string) (ip string, port int32, err error) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", -1, fmt.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	c := naming.NewNamingServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.FindService(ctx, &naming.FindServiceReq{
		ServiceName: serviceName,
	})
	if err != nil {
		return "", -1, fmt.Errorf("could not greet: %v", err)
	}
	log.Printf("rsp: %v", r)
	return r.IP, r.Port, nil
}
