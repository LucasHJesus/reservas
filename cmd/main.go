package main

import (
	"log"

	"github.com/LucasHJesus/reservas/cmd/api"
	"github.com/LucasHJesus/reservas/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewSQLStorage(mysql.Config{
		User:                 "rooot",
		Passwd:               "asd",
		Addr:                 "127.0.1:3300",
		DBName:               "reservas",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
