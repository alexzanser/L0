package subscribe

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"github.com/alexzanser/L0.git/internal/repository"
	"github.com/alexzanser/L0.git/internal/domain"
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

func Subscribe(sc stan.Conn, repo *repository.Repository) (stan.Subscription, error) {
	ReceiveMsg := func(m *stan.Msg) {
		order := &order.Order{}
		err := json.Unmarshal(m.Data, order)
		if err != nil {
			log.Println(fmt.Errorf("Cant`t unmarshal to json (invalid data)%v", err))
		} else if err := repo.Save(order.OrderUid, m.Data); err != nil {
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
