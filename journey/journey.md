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



