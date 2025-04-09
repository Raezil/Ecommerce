package services

import (
	"db"
	"generated"
	"io"
	"log"

	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

// ChatHub manages WebSocket clients and broadcasts chat messages.
type ChatHub struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan *generated.ChatMessage
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
}

// NewChatHub creates and initializes a new ChatHub.
func NewChatHub() *ChatHub {
	return &ChatHub{
		clients:    make(map[*websocket.Conn]bool),
		broadcast:  make(chan *generated.ChatMessage),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
	}
}

// Run starts the hub’s loop to handle client registration, unregistration, and broadcasting.
func (hub *ChatHub) Run() {
	for {
		select {
		case conn := <-hub.register:
			hub.clients[conn] = true
			log.Println("New WS client registered")
		case conn := <-hub.unregister:
			if _, ok := hub.clients[conn]; ok {
				delete(hub.clients, conn)
				conn.Close()
				log.Println("WS client unregistered")
			}
		case message := <-hub.broadcast:
			// Broadcast message to every registered client.
			for conn := range hub.clients {
				if err := conn.WriteJSON(message); err != nil {
					log.Printf("Error sending message via WS: %v", err)
					conn.Close()
					delete(hub.clients, conn)
				}
			}
		}
	}
}

// ChatServiceServer implements your gRPC-based chat service.
type ChatServiceServer struct {
	generated.UnimplementedChatServer
	PrismaClient *db.PrismaClient
	Logger       *zap.SugaredLogger
	hub          *ChatHub
}

// NewChatServiceServer returns a new ChatServiceServer instance.
func NewChatServiceServer(hub *ChatHub, client *db.PrismaClient, logger *zap.SugaredLogger) *ChatServiceServer {
	return &ChatServiceServer{
		PrismaClient: client,
		Logger:       logger,
		hub:          hub,
	}
}

// ChatStream implements the gRPC bidirectional streaming method.
func (s *ChatServiceServer) ChatStream(stream generated.Chat_ChatStreamServer) error {
	for {
		// Read a message from the gRPC stream.
		msg, err := stream.Recv()
		if err == io.EOF {
			// Client closed the stream.
			return nil
		}
		if err != nil {
			return err
		}

		// Log and broadcast the incoming message.
		log.Printf("gRPC Received message: %v", msg)
		s.hub.broadcast <- msg

		// Optionally, send the same message back to the gRPC client.
		if err := stream.Send(msg); err != nil {
			return err
		}
	}
}

// -------------------------------------------------------------------------
// fasthttp WebSocket Handler Integration
// -------------------------------------------------------------------------

// fastHTTPUpgrader is configured to upgrade fasthttp requests to WebSocket connections.
// In production, adjust CheckOrigin to validate incoming origins.
var fastHTTPUpgrader = websocket.FastHTTPUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(ctx *fasthttp.RequestCtx) bool {
		return true
	},
}

// ServeWs upgrades a fasthttp request to a WebSocket connection,
// registers the connection with the ChatHub, and handles incoming messages.
func (hub *ChatHub) ServeWs(ctx *fasthttp.RequestCtx) {
	err := fastHTTPUpgrader.Upgrade(ctx, func(conn *websocket.Conn) {
		// Register the new WebSocket client.
		hub.register <- conn

		// Ensure the client is unregistered when done.
		defer func() {
			hub.unregister <- conn
			conn.Close()
		}()

		// Continuously read messages from the WebSocket connection.
		for {
			var msg generated.ChatMessage
			if err := conn.ReadJSON(&msg); err != nil {
				log.Printf("WS read error: %v", err)
				break
			}

			log.Printf("WS Received message: %v", msg)
			// Broadcast the received message to all clients.
			hub.broadcast <- &msg
		}
	})

	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		ctx.Error("WebSocket upgrade failed", fasthttp.StatusBadRequest)
	}
}
