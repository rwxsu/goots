package network

import (
	"net"
	"sync"

	"github.com/rwxsu/goot/game"
)

// ConnectionManager should be used to management of connections
type ConnectionManager struct {
	connections []*TibiaConnection
	lock        *sync.Mutex
}

func NewConnectionManager() ConnectionManager {
	return ConnectionManager{
		lock: &sync.Mutex{},
	}
}

// Add should be used to add TibiaConnection to ConnectionManager
func (connectionManager *ConnectionManager) Add(tibiaConnection *TibiaConnection) {
	connectionManager.lock.Lock()
	connectionManager.connections = append(connectionManager.connections, tibiaConnection)
	connectionManager.lock.Unlock()
}

// GetByConn retrieve TibiaConnection by net.Conn
func (connectionManager *ConnectionManager) GetByConn(conn net.Conn) *TibiaConnection {
	for _, current := range connectionManager.connections {
		if current.Connection == conn {
			return current
		}
	}

	return nil
}

// Del should be used to remove TibiaConnection from ConnectionManager
func (connectionManager *ConnectionManager) Del(tibiaConnection *TibiaConnection) {
	connectionManager.lock.Lock()

	for i, current := range connectionManager.connections {
		if current == tibiaConnection {
			connectionManager.connections = append(connectionManager.connections[:i], connectionManager.connections[i+1:]...)
			break
		}
	}

	connectionManager.lock.Unlock()
}

// TibiaConnection is a struct to store pair of connection and player
type TibiaConnection struct {
	Connection net.Conn
	Player     *game.Creature
	Map        *game.Map
}
