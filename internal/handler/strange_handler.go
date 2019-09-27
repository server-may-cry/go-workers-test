package handler

import (
	"encoding/json"
	"github.com/server-may-cry/go-workers-test/internal/entity"
	"github.com/server-may-cry/go-workers-test/internal/services/demoservice"
	"net/http"
)

func New(service *demoservice.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = service.InsertObject(&entity.MyEntry{
			Message: "handler message",
			Source:  "strange handler",
		})

		lastInserted, _ := service.ReadObject()
		if err := json.NewEncoder(w).Encode(lastInserted); err != nil {
			panic(err)
		}
	}
}
