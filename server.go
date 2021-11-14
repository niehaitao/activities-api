package main

import (
	"activities-api/db"
	"activities-api/web"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	d, err := sql.Open("postgres", dsn())
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()

	app := web.NewApp(db.NewDB(d))
	err = app.Serve()
	log.Println("Error", err)
}

func dsn() string {

	host := os.Getenv("db_host")
	user := os.Getenv("db_user")
	pass := os.Getenv("db_pass")
	port := os.Getenv("db_port")
	name := os.Getenv("db_name")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		pass,
		name)

	return dsn
	// "postgresql://" + host + ":" + port + "/" + name +
	// 	"?user=" + user +
	// 	"&sslmode=disable&password=" + pass
}
