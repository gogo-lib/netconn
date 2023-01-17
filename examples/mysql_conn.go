package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gogo-lib/netconn"
)

func main() {
	// init user db connection
	pgDSN := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"192.168.49.2", "30015", "admin", "admin", "netconn-mysql",
	)

	for {
		time.Sleep(time.Second * 1)
		conn := netconn.GetPostgreSQLConn(pgDSN)
		log.Println("conn: ", conn.Ping())
	}
}
