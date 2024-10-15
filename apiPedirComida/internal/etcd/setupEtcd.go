package etcd

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/client/v3"
)

type EtcdClient struct {
	cli *clientv3.Client
}

func NewEtcdClient() *EtcdClient {
	//Create client etc
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: 20 * time.Second,
	})

	if err != nil {
		log.Fatal(err)
	}
	return &EtcdClient{cli}
}

func (etd *EtcdClient) RegisterApi(serviceName string, serverPort string) error {
	//Register service
	serviceValue := "http://localhost/" + serverPort
	ctx := context.Background()
	_, err := etd.cli.Put(ctx, serviceName, serviceValue)

	log.Println("Service registered !!")
	return err
}

func (etcd *EtcdClient) DiscoveryService() {
	ctx := context.Background()
	resp, err := etcd.cli.Get(ctx, "/services/", clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}

	if len(resp.Kvs) == 0 {
		log.Println("No service found")
		return
	}

	log.Println("List of servies:")
	for _, ev := range resp.Kvs {
		fmt.Println(string(ev.Key), string(ev.Value))
	}

}

func (etcd *EtcdClient) Close() error {
	return etcd.cli.Close()
}
