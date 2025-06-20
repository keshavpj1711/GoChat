import { useState } from "react";
import { sendMsg } from "../api";

const Connection = ({estbConn}) => {
  const [input, setInput] = useState("")
  const [inputValue, setInputValue] = useState("")

  async function ChatInput(e) {
    // console.log(input)

    // using from App.jsx
    await estbConn()

    // Now here we can open the websocket connection
    sendMsg(input)

    // Erasing the input field
    setInputValue("")
  }

  const handleKeyDown = e => {
    if (e.key  === 'Enter') {
      ChatInput()
    }
  }

  function handleChange(e) {
    // getting the input in the input tag and setting the state as it's value
    setInput(e.target.value)
    setInputValue(e.target.value)
  }

  return (
    <div className="my-2">
      <div className="px-2 flex gap-2 justify-center">
        <input
          type="text"
          placeholder="Enter text"
          name="msg"
          value={inputValue}
          onChange={handleChange}
          onKeyDown={handleKeyDown}
          className="bg-amber-50 rounded-md p-1 grow" />

        <button onClick={ChatInput}
          className="bg-celestialBlue-50 p-2 rounded-md text-amber-50 hover:scale-105 transition-all">
          Send Msg
        </button>
      </div>
    </div>
  );
}

export default Connection;