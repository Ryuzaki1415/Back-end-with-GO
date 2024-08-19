package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Ryuzaki1415/ecom/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct{   
	addr string
	db *sql.DB
}

// we need to initialize new  API server


func NewAPIServer(addr string, db *sql.DB)*APIServer{
	return &APIServer{  // returns a datatype of type APISERVER defined above.  Function to initialize a http server
		addr: addr,
		db: db,
	}

}


func (s *APIServer) Run()error{    
	router:=mux.NewRouter()     // first step after we call the run method in main, it creates a router.
	subrouter:=router.PathPrefix("/api/v1").Subrouter()  // subrouters are generally a good practice for versioning the API
	log.Println("Listening on PORT ",s.addr)
	userHandler:=user.NewHandler()   
	userHandler.RegisterRoutes(subrouter)

	return http.ListenAndServe(s.addr,router)

}


//to create function  run to run the http server


