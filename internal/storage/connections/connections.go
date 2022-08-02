package connections

import (
	"context"
	"github.com/pkg/errors"
	"time"
)

var ErrTimeout = errors.New("timeout")

type Connections struct {
	// connectionPool limits connections count
	connectionPool chan struct{}

	timeout time.Duration
}

func NewPool(poolSize int, timeout time.Duration) *Connections {
	return &Connections{
		connectionPool: make(chan struct{}, poolSize),
		timeout:        timeout,
	}
}

func (c *Connections) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	select {
	case c.connectionPool <- struct{}{}:
		return nil
	case <-ctx.Done():
		return errors.Wrap(ErrTimeout, ctx.Err().Error())
	}
}

func (c *Connections) Disconnect() {
	<-c.connectionPool
}
