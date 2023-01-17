package netconn

import (
	"database/sql"
	"sync"

	// postgres driver
	_ "github.com/lib/pq"
)

var (
	// postgreSQLConn manages all mysql connection by datasource name
	postgreSQLConn    = make(map[string]*sql.DB)
	rwmPostgreSQLConn sync.RWMutex
)

// GetPostgreSQLConn return connection to mysql
func GetPostgreSQLConn(dsn string) *sql.DB {
	// check conn is already connected
	rwmPostgreSQLConn.RLock()
	conn, ok := postgreSQLConn[dsn]
	if ok && conn != nil && conn.Ping() == nil {
		defer rwmPostgreSQLConn.RUnlock()
		return conn
	}
	rwmPostgreSQLConn.RUnlock()

	// create new connection
	loggerObj.print("netconn.GetPostgreSQLConn create new connection")

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		loggerObj.printf("netconn.GetMysqlConn %v", err)
		return conn
	}

	if pingErr := conn.Ping(); pingErr != nil {
		loggerObj.printf("netconn.GetPostgreSQLConn %v", pingErr)
		return conn
	}

	// store conn instance to storage manager
	rwmPostgreSQLConn.Lock()
	defer rwmPostgreSQLConn.Unlock()
	postgreSQLConn[dsn] = conn

	return postgreSQLConn[dsn]
}
