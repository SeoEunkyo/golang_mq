package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func ServeAPI(listenAddr string){
	r := mux.NewRouter()
	r.Methods("get").Path("/").Handler(&IndexHandler{})
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