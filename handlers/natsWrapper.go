package handlers

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/nats-io/nats.go"
	"log"
	"nats-template/connections"
	"nats-template/core/models"
	"nats-template/lib/topics"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type NATSWrapper struct {
	nats *nats.EncodedConn
}

type StringNatsHandlers interface {
	countChars(request models.CountCharsRequest) models.CountCharsResponse
	reverseString(request models.ReverseStringRequest) models.ReverseStringResponse
}

func (n *NATSWrapper) init() {
	natsConnection := connections.NATS{}
	err := natsConnection.Init()
	if err != nil {
		log.Fatal("NATS Init failed")
	}
	n.nats = natsConnection.Client
}

func (n *NATSWrapper) Start(stringNatsHandlers StringNatsHandlers) error {
	n.init()

	_, err := n.nats.Subscribe(topics.StringRequest, func(m *nats.Msg) {
		switch m.Subject {
		case topics.StringRequestCountChars:
			var countCharsRequest models.CountCharsRequest
			err := json.Unmarshal(m.Data, &countCharsRequest)
			if err != nil {
				log.Println("Invalid data Unmarshal failed: ", err)
			} else {
				go func() {
					countCharsResponse := stringNatsHandlers.countChars(countCharsRequest)
					err := n.nats.Publish(topics.StringResponseCountChars, countCharsResponse)
					if err != nil {
						log.Println(err)
					}
				}()
			}
		case topics.StringRequestReverseString:
			var reverseStringRequest models.ReverseStringRequest
			err := json.Unmarshal(m.Data, &reverseStringRequest)
			if err != nil {
				log.Println("Invalid data Unmarshal failed: ", err)
			} else {
				go func() {
					reverseStringResponse := stringNatsHandlers.reverseString(reverseStringRequest)
					err := n.nats.Publish(topics.StringResponseReverseString, reverseStringResponse)
					if err != nil {
						log.Println(err)
					}
				}()
			}
		}
	})
	if err != nil {
		log.Println("Subscription failed")
		return err
	}
	return nil
}

func (n *NATSWrapper) Close() {
	n.nats.Close()
}
