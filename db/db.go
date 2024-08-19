package db

import (
	"database/sql"
	"log"
	"github.com/go-sql-driver/mysql"
)


func NewMySQLStorage(cfg mysql.Config) (*sql.DB,error){


	db,err:=sql.Open("mysql",cfg.FormatDSN())
	if err!=nil{
		log.Fatal("THE DATABASE FAILED TO CONNECT ",err)
	}

	return db , nil
}



