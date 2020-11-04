package connections

import (
	"errors"
	"github.com/nats-io/nats.go"
	"log"
	"os"
)

type NATS struct {
	Client *nats.EncodedConn
}

func (n *NATS) Init() error {
	nc, err := nats.Connect(os.Getenv("NATS_SERVER"))
	if err != nil {
		log.Println("NATS Init failed")
		log.Fatal(err)
	}
	natsClient, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Println("NATS Init failed")
		log.Fatal(err)
	}
	n.Client = natsClient
	return nil
}

func (n *NATS) Close() error {
	if n.Client.Conn.Status() == nats.CONNECTED ||
		n.Client.Conn.Status() == nats.CONNECTING ||
		n.Client.Conn.Status() == nats.RECONNECTING {
		n.Client.Close()
		return nil
	} else {
		return errors.New("nats connection already closed")
	}
}
