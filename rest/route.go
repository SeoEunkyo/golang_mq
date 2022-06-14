package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SeoEunkyo/golang_mq/lib/msgqueue"
	"github.com/gorilla/mux"
)

func ServeAPI(listenAddr string,  eventEmitter msgqueue.EventEmitter ){
	
	
	r := mux.NewRouter()
	r.Methods("get").Path("/").Handler(&IndexHandler{eventEmitter})
	r.Methods("get").Path("/event/{eventID}/booking").Handler(&CreateBookingHandler{})

	srv := http.Server{
		Handler:      r,
		Addr:         listenAddr,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  1 * time.Second,
	}
	
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("err : " , err)
	}
}