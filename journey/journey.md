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

## 