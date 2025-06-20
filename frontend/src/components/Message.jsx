import { useEffect, useState } from "react";

const Message = ({message}) => {
  let temp = JSON.parse(message)
  console.log(temp.body)
  const [msgBody, setMsgBody] = useState("")
  
  useEffect (() => {
    setMsgBody(temp.body)
  })

  return ( 
    <div className="text-white">
      {"> " + msgBody}
    </div>
   );
}
 
export default Message;