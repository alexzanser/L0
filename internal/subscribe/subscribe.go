package subscribe

import (
	"fmt"
	"time"
	"github.com/alexzanser/L0.git/internal/store"
	stan "github.com/nats-io/stan.go"
)

const 	durableName = "waiting-for-the-end"

func Connect(clusterID, clientID string) (stan.Conn, error) {
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		return nil, fmt.Errorf("Error while Connect to STAN: %v", err)
	}
	return sc, nil 
}

func Subscribe(sc stan.Conn, store store.Storage) (stan.Subscription, error) {
	ReceiveMsg := func(m *stan.Msg) {
		store.Save(m.Data)
		m.Ack()
	}

	aw, _ := time.ParseDuration("60s")
	sub, err := sc.Subscribe("foo", 
				ReceiveMsg, 
				stan.SetManualAckMode(), 
				stan.AckWait(aw), 
				stan.DurableName(durableName))
	if err != nil {
		return nil, fmt.Errorf("Error subcribing to channel: %v", err)
	}
	return sub, nil
}
