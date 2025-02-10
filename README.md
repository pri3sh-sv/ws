# WebSocket Server in Go using Gorilla

## Overview
This project is a structured WebSocket server built with Go and the Gorilla WebSocket package. It provides a clean architecture with separate modules for handling WebSocket connections, client interactions, and message broadcasting.

## Features
- **Modular design** with separate packages for WebSocket logic.
- **Broadcasting support** to send messages to all connected clients.
- **Graceful handling of connections** (register, unregister, and message handling).
- **Ping/Pong mechanism** to keep connections alive.
- **Structured logging** for debugging purposes.

## Project Structure
```
/websocket-server
│── /cmd
│    └── main.go            # Entry point for the application
│── /internal
│    ├── /server
│    │   ├── server.go      # Initializes the HTTP & WebSocket server
│    ├── /ws
│    │   ├── client.go      # Manages individual WebSocket clients
│    │   ├── hub.go         # Manages WebSocket connections and broadcasts
│    │   ├── handler.go     # Handles WebSocket HTTP upgrades
│── go.mod
│── go.sum
│── README.md
```

## Installation
1. **Clone the Repository**
```sh
git clone https://github.com/your-username/websocket-server.git
cd websocket-server
```
2. **Install Dependencies**
```sh
go mod tidy
```

## Running the WebSocket Server
```sh
go run cmd/main.go
```
The server will start on `http://localhost:8080/ws`.

## WebSocket API
- **Endpoint:** `ws://localhost:8080/ws`
- **Methods:**
    - Clients can send messages to the server.
    - The server broadcasts messages to all connected clients.

## Testing WebSocket Connection
### Using wscat (WebSocket Client)
```sh
npm install -g wscat
wscat -c ws://localhost:8080/ws
```

### Using HTML & JavaScript Client
Create an `index.html` file:
```html
<!DOCTYPE html>
<html>
<head>
    <title>WebSocket Test</title>
</head>
<body>
    <script>
        let ws = new WebSocket("ws://localhost:8080/ws");

        ws.onopen = () => {
            console.log("Connected");
            ws.send("Hello Server!");
        };

        ws.onmessage = (event) => {
            console.log("Received:", event.data);
        };

        ws.onclose = () => {
            console.log("Disconnected");
        };
    </script>
</body>
</html>
```

Open this file in a browser and check the console for WebSocket activity.


