package main

import (
	"fmt"

	"github.com/gogo-lib/netconn"
)

func main() {
	// init user db connection
	pgDSN := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"192.168.49.2", "30015", "admin", "admin", "netconn-mysql",
	)

	conn := netconn.GetPostgreSQLConn(pgDSN)
	fmt.Println(conn.Ping())
}
