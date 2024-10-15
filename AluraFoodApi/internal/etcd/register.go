package etcd

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/client/v3"
)

func RegisterApi(serviceName string, serverPort string) {

	//Create etcd client
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: 20 * time.Second,
	})

	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	//Register service
	serviceValue := "http://localhost/" + serverPort
	_, err = cli.Put(context.Background(), serviceName, serviceValue)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Service registered in etcd!!")

}

func DiscoverApi() {
	//Create etcd client
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: 20 * time.Second,
	})

	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	rest, err := cli.Get(context.Background(), "/services/", clientv3.WithPrefix())

	if err != nil {
		log.Fatal(err)
	}

	for _, r := range rest.Kvs {
		fmt.Println(string(r.Key), string(r.Value))
	}

}
