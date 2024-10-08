package monitor

import (
	"log/slog"
	"os"
	"time"

	"github.com/cloudfoundry/dropsonde/metrics"

	log "code.cloudfoundry.org/gorouter/logger"
)

//go:generate counterfeiter -o ../fakes/fake_subscriber.go . Subscriber
type Subscriber interface {
	Pending() (int, error)
	Dropped() (int, error)
}

type NATSMonitor struct {
	Subscriber Subscriber
	Sender     metrics.MetricSender
	TickChan   <-chan time.Time
	Logger     *slog.Logger
}

func (n *NATSMonitor) Run(signals <-chan os.Signal, ready chan<- struct{}) error {
	close(ready)
	for {
		select {
		case <-n.TickChan:
			queuedMsgs, err := n.Subscriber.Pending()
			if err != nil {
				n.Logger.Error("error-retrieving-nats-subscription-pending-messages", log.ErrAttr(err))
			}
			chainer := n.Sender.Value("buffered_messages", float64(queuedMsgs), "message")
			err = chainer.Send()
			if err != nil {
				n.Logger.Error("error-sending-buffered-messages-metric", log.ErrAttr(err))
			}

			droppedMsgs, err := n.Subscriber.Dropped()
			if err != nil {
				n.Logger.Error("error-retrieving-nats-subscription-dropped-messages", log.ErrAttr(err))
			}
			chainer = n.Sender.Value("total_dropped_messages", float64(droppedMsgs), "message")
			err = chainer.Send()
			if err != nil {
				n.Logger.Error("error-sending-total-dropped-messages-metric", log.ErrAttr(err))
			}
		case <-signals:
			n.Logger.Info("exited")
			return nil
		}
	}
}
