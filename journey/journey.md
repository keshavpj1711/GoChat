# Setting up the Frontend and Backend

For Frontend I am using ReactJS and for Backend I am using GoLang.

## Setting up Frontend

- To create a basic layout of our react app - `npx create-react-app dir_name`
  - here **dir_name** is **frontend**
- Now we have our starting files for our react app

## Setting up Backend

- Create a directory named backend in the same level as frontend, and then move into the dir - `mkdir backend && cd backend`
- To initialize go modules we do - `go mod init module_path`
- Now finally check if the go program is running or not.

# Setting up communication

In this part we will be implementing a basic WebSocket server which will listen for messages and write them back to via the same websocket.

## Setting up backend WebSocket server

What are we using: **Gorilla web toolkit**
- This toolkit provides useful packages for writing HTTP-based applications.
- In this we will be using **gorilla/websockets** to setup our websocket server

Adding it as dependency
```bash
go get github.com/gorilla/websocket
```

## Web Socket Protocol 

- WebSockets offer us duplex communication from a non trusted to a server that we own across a TCP socket connection. 
- This means that instead of continuously polling our server for updates and having to perform TCP handshakes every time we poll, we can maintain a single TCP socket connection and then send and listen to message on that.

> This reduces the network overhead by a huge amount, which is the main deal for any real time application and helps maintain incredible amount of clients on a single server instance.

### The Cons ⚠️

- As soons as you introduce state(i.e. memory of each connection), it becomes more complex with regards to scaling up your app across multiple instances.
  - Why? because now the server needs to remember things about a specific connection.

> Now scaling up means you **run multiple copies of your app** so when your app grows you usually run many copies of it on different servers or containers

- So in order to solve this issue that comes with scaling up, we consider options like storing your state in message brokers(like Redis Pub/Sub, Kafka), or in db/mem caches(Redis or Memcached) that can scale in parallel with your application instances.
  - This adds complexity.

## Implementation 

First we create a new endpoint and then upgrade the connection to long lasting WebSocket connection.

Thankfully, the gorilla/websocket package features the functionality for upgrading a HTTP connection to a WebSocket connection.

### Creating a WebSocket Endpoint

Creating a new endpoint `/ws` and converting from a std `http` endpoint to `/ws`.

This endpoint will do 3 things: 
- Check the origin of incoming HTTP request
  - return `true` for every client
    - since as of now we open to each client.
- Attempt to upgrade the connection using a defined `upgrader`
- Listen and Echo messages: so once the websocket conn is open
  - the server waits for msgs from clients
  - prints the msg to logs, for us to see
  - echo the msg back to same conn
  > helps to test the connection works both ways

<details>
<summary>Code Snippets</summary>
  
### Websocket upgrader

```go
// This is like the handshake in WebSockets, and
// in Go we specifically use this to turn a normal HTTP req to a websocket conn
var upgrader = websocket.Upgrader{ 
  // Reading and writing buffer sizes
  ReadBufferSize: 1024,
  WriteBufferSize: 1024,
  
  // Allow all client to connect for now
  // Similar to CORS in Node.js -- how you'd allow react(localhost:5000 in this case) to talk to your backend
  CheckOrigin: func(r *http.Request) bool { return true },
}

```

### To handle the websocket connection

```go
// The function that handles requests to /ws route
func serveWs(w http.ResponseWriter, r *http.Request) {
  // this is where we finally use tht upgrader to upgrade the conn
  ws, err := upgrader.Upgrade(w, r, nil)
  // Handles the error

  // start listening for messages on the websocket conn using reader
  reader(ws)
}
```

### Reader: one that reads every msg to ws conn and echos it back

```go
func Reader(conn *websocket.Conn) {
  for {
    // Read a message from the WebSocket
    messageType, p, err := conn.ReadMessage()
    if err != nil {
      log.Println(err)
      return
    }
    
    // Print the message we received
    fmt.Println(string(p))

    // Echo the message back to the same connection
    if err := conn.WriteMessage(messageType, p); err != nil {
      log.Println(err)
      return
    }
  }
}
```

</details>


### Creating the React WebSocket Client

This Websocket client handles the connection request to the socket and sending the messages to the ws conn.

```js
var socket = new WebSocket("ws://localhost:3000/ws")

let connect = cb => {
  console.log("Attempting Connection...")

  socket.onopen = () => {
    console.log("Successfully Connected")
  }

  socket.onmessage = msg => {
    console.log(msg);
    cb(msg)
  }

  socket.onclose = e => {
    console.log("Socket Closed Connection: ", e)
  }

  socket.onerror = err => {
    console.log("Socket Error: ", err)
  }
}

let sendMsg = msg => {
  console.log("Sending Msg: ", msg)
  socket.send(msg)
}

export {
  connect, 
  sendMsg
}
```

## Handling Multiple Clients

In simple terms what we need to do is whenever a new connection is made, we'll have to add them to a pool of existing connections and ensure that every time a message is sent, everyone in that pool recieves that message.

### Using Channels 

- For each of these concurrent connections a **new `goroutine` is spun up** for that duration of that connection. 

- This means that we have to worry about communication across these concurrent `goroutines` and ensure that whatever we are doing, is thread-safe. 

- So, when we are implementing a **Pool** structure further down the line, we’ll have to consider either using a `sync.Mutex` to mutually exclude `goroutines` from simultaneously accessing/modifying our data, or we can use `channels`.

> Here i think we'll be better off using `channels` and using them to communicate in a safe fashion across these multiple concurrent `goroutines`, since i am most comfortable with that 😅


