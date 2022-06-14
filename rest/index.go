package rest

import (
	"net/http"

	"github.com/SeoEunkyo/golang_mq/contracts"
	"github.com/SeoEunkyo/golang_mq/lib/msgqueue"
)

type IndexHandler struct {
	eventEmitter msgqueue.EventEmitter
}
func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	// MQ 생성하기
	msg := contracts.EventBookedEvent{
		EventID: "1",
		UserID:  "someUserID",
	}
	h.eventEmitter.Emit(&msg)
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write( []byte("Server is running "))
}