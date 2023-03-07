package dgraph

import (
	"log"
	"sync"

	//"os"
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

type Client struct {
	dg *dgo.Dgraph
}

var once sync.Once
var client *Client

func initDg() *dgo.Dgraph {
	//conn, err := grpc.Dial(os.Getenv("DG_SERVER_ANALYTICS"), grpc.WithInsecure(), grpc.WithMaxMsgSize(4 << 30))
	conn, err := grpc.Dial("dg-info-grpc.fouita.com:80", grpc.WithInsecure(), grpc.WithMaxMsgSize(4<<30))
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)
	return dg
}

func Server() *Client {
	once.Do(func() {
		client = &Client{dg: initDg()}
	})
	return client
}

func (*Client) Get() *dgo.Dgraph {
	return client.dg
}
