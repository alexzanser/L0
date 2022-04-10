package subscribe

import (
	"fmt"
	"log"
	"time"

	"github.com/alexzanser/L0.git/internal/repository"
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

func Subscribe(sc stan.Conn, store *repository.Storage) (stan.Subscription, error) {
	ReceiveMsg := func(m *stan.Msg) {
		if err := store.Save(m.Data); err != nil {
			log.Println(err)
		}
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
