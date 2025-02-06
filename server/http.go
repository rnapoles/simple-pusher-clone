package server

import (
	"simple-pusher-clone/utils"
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Global variable for the socket.io server
var server *socketio.Server

// Initialize the socket.io server
func initSocketIO() *socketio.Server {

	server = socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})

	// Handle new connections
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("New connection: ", s.ID())
		return nil
	})

	// Handle joining rooms (channels)
	server.OnEvent("/", "join", func(s socketio.Conn, channel string) {
		s.Join(channel)
		log.Printf("User %s joined channel %s\n", s.ID(), channel)
	})

	// Handle disconnection
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		server.LeaveAllRooms("/", s)
		log.Printf("User %s disconnected: %s\n", s.ID(), reason)
	})

	// Handle errors
	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("Error: ", e)
	})

	return server
}

func allowOriginFunc(r *http.Request) bool {
	return true
}

// REST endpoint to send a message to a specific user
func sendMessageToUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	message := r.URL.Query().Get("message")

	if userID == "" || message == "" {
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		return
	}

	// Emit a message to a specific user by their connection ID
	server.ForEach("/", userID, func(c socketio.Conn) {
		c.Emit("message", message)
	})

	w.Write([]byte("Message sent to user: " + userID))
}

// REST endpoint to send a message to a specific channel
func sendMessageToChannel(w http.ResponseWriter, r *http.Request) {
	// Retrieve the API key from the request header
	/*apiKey := r.Header.Get("X-API-KEY")

	// Check if the API key is provided
	if apiKey == "" {
		server.Error(w, "Missing API key", server.StatusUnauthorized)
		return
	}

	// Validate the API key (replace "your-secret-api-key" with your actual key)
	if apiKey != "your-secret-api-key" {
		server.Error(w, "Invalid API key", server.StatusForbidden)
		return
	}*/

	// Proceed with the existing logic if the API key is valid
	vars := mux.Vars(r)
	channelID := vars["channelID"]
	//message := r.URL.Query().Get("message")

	if channelID == "" /* || message == ""*/ {
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		return
	}

	// Broadcast a message to all users in a specific channel
	server.BroadcastToRoom("/", channelID, "message", channelID)

	// Log
	log.Printf("Notify channel: %s\n", channelID)

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Respond with a JSON object indicating success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": true}`))
}

func StartServer() {

	listen := utils.GetEnv("LISTEN", "127.0.0.1")
	port := utils.GetEnv("PORT", "9002")
	address := fmt.Sprintf("%s:%s", listen, port)

	// Create a new router
	router := mux.NewRouter()

	// Initialize the socket.io server
	socketIOServer := initSocketIO()

	// Serve the socket.io WebSocket endpoint
	router.Handle("/socket.io/", socketIOServer)

	// Add REST endpoints for sending messages
	router.HandleFunc("/send/user/{userID}", sendMessageToUser).Methods(http.MethodGet)
	router.HandleFunc("/send/channel/{channelID}", sendMessageToChannel).Methods(http.MethodGet)

	// Serve the static HTML and other files from the "static" folder
	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))

	// Handle invalid (404) routes
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 - Route Not Found", http.StatusNotFound)
		log.Printf("Route not found: %s, Methods: %v\n", r.Method, r.URL)
	})

	err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		log.Printf("Route: %s, Methods: %v\n", pathTemplate, methods)
		return nil
	})

	if err != nil {
		return
	}

	// Start the WebSocket server in a goroutine
	go func() {
		if err := socketIOServer.Serve(); err != nil {
			log.Fatalf("SocketIO serve error: %s\n", err)
		}
	}()
	defer socketIOServer.Close()

	// Start the HTTP server
	log.Printf("Server started on %s", address)
	log.Fatal(http.ListenAndServe(address, router))
}
