# Realtime Interactive Server using Socket.IO and Golang

## Overview
This project is a lightweight alternative to Pusher, built with Golang and Socket.IO. It enables developers to create powerful real-time interactive applications, handling events, messages, and live updates efficiently.

## Features
- WebSocket-based real-time communication
- Event-driven architecture
- Broadcast and private messaging support
- Channel-based communication like Pusher
- Simple integration with frontend applications
- Scalable and lightweight

## Prerequisites
- Golang (>=1.18)
- Node.js (for frontend testing)

## Installation

Clone the repository:
```sh
git clone simple-pusher-clone.git
cd simple-pusher-clone
```

Install dependencies:
```sh
go mod tidy
```

## Usage

Run the server:
```sh
go run main.go
```

The server will start on `http://localhost:9002` by default.

### Example Frontend Connection
To connect to the server using JavaScript:
```js
const socket = io('http://localhost:9002',{
  autoConnect: true,
  path: '/socket.io',
  transports: ['websocket'],
});
    
socket.on("connect", () => {
    // Join a channel
    socket.emit("join", "channel1");

    // Listen for messages
    socket.on("message", function(data) {
        console.log("Received message: " + data);
    });
});

socket.on('message', (data) => {
    console.log('Received:', data);
});
```

### API Endpoints
#### `GET /send/user/{userID}?message={message}`
Send a message to a specific user.

Example:
```sh
curl -X GET "http://localhost:8080/send/user/123?message=Hello, User!"
```

#### `GET /send/channel/{channelID}?message={message}`
Broadcast a message to all users in a specific channel.

Example:
```sh
curl -X GET "http://localhost:8080/send/channel/general?message=Hello, Channel!"
```

## Deployment
Build the application:
```sh
go build -o simple-pusher-clone
```
Run the compiled binary:
```sh
./simple-pusher-clone
```

## License
This project is licensed under the MIT License.

## Contributing
Feel free to open issues and submit pull requests to improve the project.


