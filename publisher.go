package main

import (
	"io"
	"os"
	"time"

	stan "github.com/nats-io/stan.go"
)

const (
	clusterID = "test-cluster"
	clientID = "client-publisher"
)

func main() {
	f, _ := os.Open("./model.json")
	data, _ := io.ReadAll(f)

	sc, _ := stan.Connect(clusterID, clientID)
	defer sc.Close()
	for {
		sc.Publish("foo", data)
		time.Sleep(time.Second * 5)
	}
}
