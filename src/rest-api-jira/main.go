package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {

	cfg := mysql.Config{
		User:                Envs.DBUser,
		Passwd:              Envs.DBPassword,
		Addr:                Envs.DBAddress,
		DBName:              Envs.DBName,
		Net:                 "tcp",
		AllowNativePassword: true,
		ParseTime:           true,
	}
	sqlStorage := NewMySQLStorage(cfg)

	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}
	store := NewStore(db)
	api := NewAPIServer(":4000", store)
	api.Serve()
}