package main //entry point for the api
import (
	"log"
	"database/sql"
	"github.com/Ryuzaki1415/ecom/cmd/api"
	"github.com/Ryuzaki1415/ecom/db"
	"github.com/go-sql-driver/mysql"
	"github.com/Ryuzaki1415/ecom/config"


)

func main() {

	
	//database connection

	db, err:= db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err!=nil{
		log.Fatal(err)
	}


	
	initStorage(db)

	//server initialization
	server := api.NewAPIServer(":8080",db)
	if err:= server.Run();err!=nil{
			log.Fatal(err)
		}
	}	

		//initializings storage

	func initStorage(db *sql.DB){
		err:=db.Ping()
		if err!=nil{
			log.Fatalln("THE DATABASE COULD NOT BE INITIALIZED ",err)
		}
		log.Println("THE DATABASE HAS BEEN SUCCESSFULLY CONNECTED!!")
	}