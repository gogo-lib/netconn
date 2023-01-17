package netconn

import (
	"errors"
	"strconv"
	"strings"
	"sync"

	"github.com/aerospike/aerospike-client-go"
)

var (
	// aerospikeConn manages all aerospike connection by hosts name
	aerospikeConn    = make(map[string]*aerospike.Client)
	rwmAerospikeConn sync.RWMutex
)

// GetAerospikeConn return an active connection to aerospike
func GetAerospikeConn(hosts string) *aerospike.Client {
	rwmAerospikeConn.RLock()
	conn, ok := aerospikeConn[hosts]
	if ok && conn != nil && conn.IsConnected() {
		defer rwmAerospikeConn.RUnlock()
		return conn
	}
	rwmAerospikeConn.RUnlock()

	policy := aerospike.NewClientPolicy()
	client, err := aerospike.NewClientWithPolicyAndHost(policy, newAerHost(hosts)...)
	if err != nil {
		loggerObj.printf("netconn.GetAerospike %v", err)
		return conn
	}
	if !client.IsConnected() {
		err := errors.New("aerospike client isn't ready to talk to the database server")
		loggerObj.printf("netconn.GetAerospike %v", err)
	}

	// store client instance to storage manager
	rwmAerospikeConn.Lock()
	aerospikeConn[hosts] = client
	rwmAerospikeConn.Unlock()

	return client
}

// conv name:port,name:port => []*aerospike.Host
func newAerHost(hosts string) []*aerospike.Host {
	listHostRaw := strings.Split(hosts, ",")
	listHost := make([]*aerospike.Host, 0)

	for i := 0; i < len(listHostRaw); i++ {
		hostRaw := strings.Split(listHostRaw[i], ":")
		name, portRaw := hostRaw[0], hostRaw[1]
		port, err := strconv.Atoi(portRaw)
		if err != nil {
			panic(err)
		}

		listHost = append(listHost, aerospike.NewHost(name, port))
	}

	return listHost
}
