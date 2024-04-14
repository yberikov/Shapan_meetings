package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func CreateConn() {
	// Connection parameters
	dbUser := "root"
	dbPass := "YuTu48hiwA"
	dbName := "shapan"
	dbHost := "10.101.34.22" // or your MySQL server IP address
	dbPort := "3306"         // MySQL default port

	// Create a connection string
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// Open a database connection
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to MySQL database!")

}
