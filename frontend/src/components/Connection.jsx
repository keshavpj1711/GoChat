import { useState } from "react";
import { connect, sendMsg } from "../api";

const Connection = ({estbConn}) => {
  const [input, setInput] = useState("")

  async function handleClick() {
    // console.log(input)

    await estbConn()

    // Now here we can open the websocket connection
    sendMsg(input)
  }

  function handleChange(e) {
    // getting the input in the input tag and setting the state as it's value
    setInput(e.target.value)
  }

  return (
    <div className="my-2">
      <div className="px-2 flex gap-2 justify-center">
        <input
          type="text"
          placeholder="Enter text"
          name="msg"
          onChange={handleChange}
          className="bg-amber-50 rounded-md p-1" />

        <button onClick={handleClick}
          className="bg-sky-500 p-2 rounded-md text-amber-50 hover:tras">
          Send Msg
        </button>
      </div>
    </div>
  );
}

export default Connection;