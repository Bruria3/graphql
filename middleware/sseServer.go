package middleware

import (
	"log"
)

// SSEClient represents an SSE client connection
type SSEClient struct {
	ID     int
	Stream chan string
}

// SSEServer manages clients and broadcasting messages
type SSEServer struct {
	Clients      map[int]*SSEClient
	AddClient    chan *SSEClient
	RemoveClient chan int
	Broadcast    chan string
}

func NewSSEServer() *SSEServer {
	return &SSEServer{
		Clients:      make(map[int]*SSEClient),
		AddClient:    make(chan *SSEClient),
		RemoveClient: make(chan int),
		Broadcast:    make(chan string),
	}
}

// Run starts the SSE server and handles adding/removing clients and broadcasting messages
func (server *SSEServer) Run() {
	for {
		select {
		case client := <-server.AddClient:
			server.Clients[client.ID] = client
			log.Printf("Client %d connected", client.ID)

		case id := <-server.RemoveClient:
			delete(server.Clients, id)
			log.Printf("Client %d disconnected", id)

		case message := <-server.Broadcast:
			for _, client := range server.Clients {
				client.Stream <- message
			}
		}
	}
}
