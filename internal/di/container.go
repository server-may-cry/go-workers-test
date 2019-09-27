package di

import (
	"github.com/server-may-cry/go-workers-test/internal/db"
	"github.com/server-may-cry/go-workers-test/internal/handler"
	"github.com/server-may-cry/go-workers-test/internal/services/demoservice"
	"github.com/server-may-cry/go-workers-test/internal/services/goroutineproducer"
	"net/http"
)

type Container struct {
	demoService    *demoservice.Service
	StrangeHandler http.HandlerFunc
	DummyWorker    *goroutineproducer.DummyWorker
}

func New() *Container {
	dbConnection := db.OpenConnection(db.ConnectionConfiguration{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "user",
		Password: "password",
		DBName:   "db_name",
	})

	myEntryRepository := db.NewMyEntryRepository(dbConnection)
	demoService := demoservice.New(myEntryRepository)

	return &Container{
		StrangeHandler: handler.New(demoService),
		DummyWorker:    goroutineproducer.New(demoService),
	}
}
