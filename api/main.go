package main

import (
	"fmt"
	"time"
	stan "github.com/nats-io/stan.go"
)

const (
	clusterID = "test-cluster"
	clientID = "client-222"
	durableName = "waiting-for-the-end"
)

func receiveMsg(m *stan.Msg) {
	m.Ack()
	fmt.Printf("Received a message: %s\n", string(m.Data))
}

func main() {
	quit := make(chan struct{})
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		fmt.Errorf("Error connecting to STAN: %v", err)
		return
	} 
	defer sc.Close()

	aw, _ := time.ParseDuration("60s")
	sub, err := sc.Subscribe("foo", receiveMsg, stan.SetManualAckMode(), stan.AckWait(aw), stan.DurableName(durableName))
	if err != nil {
		fmt.Errorf("Error subcribing to channel: %v", err)
		return
	}
	defer sub.Unsubscribe()
	<-quit
}
