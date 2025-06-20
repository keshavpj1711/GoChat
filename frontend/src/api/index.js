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