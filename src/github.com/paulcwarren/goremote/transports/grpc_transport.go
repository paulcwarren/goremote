// Another im
package transports

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/paulcwarren/goremote"
)

type GrpcTransport struct {
	Target string
}

func NewGrpcTransport(target string) goremote.Transport {
	return &GrpcTransport{
			Target: target,
		}
}

func (t *GrpcTransport) Send(tgt string, message interface{}, reply interface{}) (error) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(t.Target, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	err = grpc.Invoke(context.Background(), "/helloworld.Greeter/SayHello", message, reply, conn)
	if err != nil {
		return err
	}
	return nil

}
