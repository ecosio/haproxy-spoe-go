package agent

import (
	"net"

	"github.com/ecosio/haproxy-spoe-go/request"
	"github.com/ecosio/haproxy-spoe-go/worker"
)

func New(handler func(*request.Request)) *Agent {
	agent := &Agent{
		handler: handler,
	}

	return agent
}

type Agent struct {
	handler func(*request.Request)
}

func (agent *Agent) Serve(listener net.Listener) error {
	for {
		conn, err := listener.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				continue
			}
			return err
		}

		go worker.Handle(conn, agent.handler)
	}
}
