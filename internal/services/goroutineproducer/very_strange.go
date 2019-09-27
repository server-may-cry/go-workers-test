package goroutineproducer

import (
	"github.com/server-may-cry/go-workers-test/internal/entity"
	"github.com/server-may-cry/go-workers-test/internal/services/demoservice"
	"log"
	"time"
)

type DummyWorker struct {
	service *demoservice.Service
}

func New(service *demoservice.Service) *DummyWorker {
	return &DummyWorker{
		service: service,
	}
}

func (w *DummyWorker) Go() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			_ = w.service.InsertObject(&entity.MyEntry{
				Message: "worker message",
				Source: "dummy worker",
			})
			readed, _ := w.service.ReadObject()
			log.Printf("dummy worker result %+v", readed)
		}
	}
}
