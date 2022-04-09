package main

import (
	"fmt"
	"time"
	stan "github.com/nats-io/stan.go"
)

const (
	clusterID = "test-cluster"
	clientID = "client-publisher"
)

func main() {
	sc, _ := stan.Connect(clusterID, clientID)
	defer sc.Close()
	for {
		sc.Publish("foo", []byte(fmt.Sprintf("Hello World : %s", time.Now().String())))
		time.Sleep(time.Second * 1)
	}
}
