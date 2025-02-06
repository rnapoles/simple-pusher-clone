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
git clone https://github.com/yourusername/realtime-socketio-go.git
cd realtime-socketio-go
```

Install dependencies:
```sh
go mod tidy
```

## Configuration
Create a `.env` file in the root directory and define the following environment variables:

```ini
PORT=9002
LISTEN=0.0.0.0
```

- `PORT`: The port on which the server will listen.
- `LISTEN`: The IP address the server will bind to (use `0.0.0.0` to accept connections from all network interfaces).

## Usage

Run the server:
```sh
go run main.go
```

The server will start on `http://localhost:8080` by default.

### Example Frontend Connection
To connect to the server using JavaScript:
```js
const socket = io('http://localhost:8080');

socket.on('connect', () => {
    console.log('Connected to server');
    socket.emit('message', 'Hello from client');
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
go build -o realtime-server
```
Run the compiled binary:
```sh
./realtime-server
```

## License
This project is licensed under the MIT License.

## Contributing
Feel free to open issues and submit pull requests to improve the project.

## Author
[Your Name] - [Your GitHub Profile]

