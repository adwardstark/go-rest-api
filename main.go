package main

import (
	"flag"
	"github.com/adwardstark/go-rest-api/app"
)

/**
 * Created by Aditya Awasthi on 29/07/2021.
 * @author github.com/adwardstark
 */

func main() {
	var db_name string
	var db_username string
	var db_password string
	var port string

	flag.StringVar(&db_name, "db_name", "products_db", "Database name [ default: products_db ]")
	flag.StringVar(&db_username, "db_user", "postgres", "Database username [ default: postgres ]")
	flag.StringVar(&db_password, "db_pass", "pq_password", "Database password [ default: pq_password ]")
	flag.StringVar(&port, "port", "8010", "Port number to run server on [ default: 8010 ]")
	flag.Parse()

	app := app.RESTApp{}
	app.Initialize(db_username, db_password, db_name)
	app.Run(port)
}
