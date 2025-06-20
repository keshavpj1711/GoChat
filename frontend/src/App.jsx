import { useState } from 'react'
import ChatContainer from './components/ChatContainer'
import Navbar from './components/Navbar'
import { connect } from './api'

function App() {
  const [chatHistory, setChatHistory] = useState([])

  // Calling the connect function in our components life-cycle
  async function componentDidMount() {
    connect((msg) => {
      console.log("New Message: ", msg)
      setChatHistory(
        prevChatHistory => [
          ...prevChatHistory, msg.data
        ]
      )
      console.log(chatHistory)
    })
  }

  return (
    <div className='h-lvh bg-prussianBlue-50 flex flex-col'>
      <Navbar />
      <ChatContainer messages={chatHistory} onEstbConn={componentDidMount}/>
    </div>
  )
}

export default App
