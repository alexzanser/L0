package main

import (
	"encoding/json"
	"fmt"
	// "fmt"
	"io"
	"log"
	"os"
	"time"

	domain "github.com/alexzanser/L0.git/internal/domain"
	stan "github.com/nats-io/stan.go"
)

const (
	clusterID = "test-cluster"
	clientID = "client-publisher"
)

func main() {
	f, _ := os.Open("task/model.json")
	data, _ := io.ReadAll(f)
	order := &domain.Order{}
	_ = json.Unmarshal(data, order)

	i := 0
	order.OrderUid = fmt.Sprintf("%s%d", order.OrderUid, i)
	sc, _ := stan.Connect(clusterID, clientID)
	defer sc.Close()
	for {
		order.OrderUid = fmt.Sprintf("%s%d", order.OrderUid[:len(order.OrderUid) - 1], i)
		data, _ = json.Marshal(order)
		sc.Publish("foo", data)
		log.Printf("%s\n", order.OrderUid)
		time.Sleep(time.Second * 5)
		i++
	}
}
