package services

import (
	"db"
	"generated"
	"io"
	"log"
	"sync"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

// -----------------------------------------------------------------------------
// ChatRoom represents an individual chat room.
type ChatRoom struct {
	name       string
	clients    map[*websocket.Conn]bool
	broadcast  chan *generated.ChatMessage
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
}

// NewChatRoom creates and initializes a new ChatRoom.
func NewChatRoom(name string) *ChatRoom {
	return &ChatRoom{
		name:       name,
		clients:    make(map[*websocket.Conn]bool),
		broadcast:  make(chan *generated.ChatMessage),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
	}
}

// Run starts the room’s loop to manage client registration, unregistration,
// and message broadcasting.
func (room *ChatRoom) Run() {
	for {
		select {
		// Process new registrations.
		case conn := <-room.register:
			room.clients[conn] = true
			log.Printf("New client registered to room %s", room.name)

		// Process unregistrations.
		case conn := <-room.unregister:
			if _, ok := room.clients[conn]; ok {
				delete(room.clients, conn)
				conn.Close()
				log.Printf("Client unregistered from room %s", room.name)
			}

		// Process incoming messages.
		case msg := <-room.broadcast:
			// Marshal the proto message using protojson.
			jsonData, err := protojson.Marshal(msg)
			if err != nil {
				log.Printf("Error marshaling message in room %s: %v", room.name, err)
				continue
			}
			// Broadcast the marshaled message to each registered client.
			for conn := range room.clients {
				if err := conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
					log.Printf("Error sending message in room %s: %v", room.name, err)
					conn.Close()
					delete(room.clients, conn)
				}
			}
		}
	}
}

// -----------------------------------------------------------------------------
// ChatHub manages multiple chat rooms.
type ChatHub struct {
	rooms map[string]*ChatRoom
	mu    sync.Mutex // protects the rooms map
}

// NewChatHub creates a new ChatHub instance.
func NewChatHub() *ChatHub {
	return &ChatHub{
		rooms: make(map[string]*ChatRoom),
	}
}

// GetRoom retrieves an existing room by name or creates a new one if it doesn't exist.
func (hub *ChatHub) GetRoom(name string) *ChatRoom {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	room, exists := hub.rooms[name]
	if !exists {
		room = NewChatRoom(name)
		hub.rooms[name] = room
		go room.Run() // Start the chat room's message loop.
		log.Printf("Created new chat room: %s", name)
	}
	return room
}

// -----------------------------------------------------------------------------
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
// Here we broadcast received messages to a default room.
func (s *ChatServiceServer) ChatStream(stream generated.Chat_ChatStreamServer) error {
	// For this example, we use a default room named "default"
	room := s.hub.GetRoom("default")

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
		room.broadcast <- msg

		// Optionally, send the same message back to the gRPC client.
		if err := stream.Send(msg); err != nil {
			return err
		}
	}
}

// -----------------------------------------------------------------------------
// FastHTTP/WebSocket Integration
// -----------------------------------------------------------------------------
// fastHTTPUpgrader is configured to upgrade fasthttp requests to WebSocket connections.
var fastHTTPUpgrader = websocket.FastHTTPUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// In production, adjust CheckOrigin to validate incoming origins appropriately.
	CheckOrigin: func(ctx *fasthttp.RequestCtx) bool {
		return true
	},
}

// ServeWs upgrades a fasthttp request to a WebSocket connection,
// uses a query parameter (e.g. ?room=myroom) to join a specific chat room,
// registers the connection with that room, and handles incoming messages.
// ServeWs upgrades a fasthttp request to a WebSocket connection,
// uses a query parameter (e.g. ?room=myroom) to join a specific chat room,
// registers the connection with that room, and handles incoming messages.
func (hub *ChatHub) ServeWs(ctx *fasthttp.RequestCtx) {
	// Extract the room name from query parameters; default to "default" if empty.
	roomName := string(ctx.QueryArgs().Peek("room"))
	if roomName == "" {
		roomName = "default"
	}

	// Get or create the chat room.
	room := hub.GetRoom(roomName)

	err := fastHTTPUpgrader.Upgrade(ctx, func(conn *websocket.Conn) {
		// Register the new client with the chosen room.
		room.register <- conn

		// Ensure the client is unregistered on exit.
		defer func() {
			room.unregister <- conn
			conn.Close()
		}()

		// Handle incoming messages from this WebSocket connection.
		for {
			// Read the raw WebSocket message.
			msgType, data, err := conn.ReadMessage()
			if err != nil {
				log.Printf("WS room %s read error: %v", room.name, err)
				break
			}
			// Process only text messages.
			if msgType != websocket.TextMessage {
				continue
			}

			// Unmarshal using the protobuf JSON unmarshaler.
			var msg generated.ChatMessage
			if err := protojson.Unmarshal(data, &msg); err != nil {
				log.Printf("protojson.Unmarshal error in room %s: %v", room.name, err)
				continue
			}

			log.Printf("WS room %s received message: %+v", room.name, msg)
			room.broadcast <- &msg
		}
	})

	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		ctx.Error("WebSocket upgrade failed", fasthttp.StatusBadRequest)
	}
}
