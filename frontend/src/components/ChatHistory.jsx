const ChatHistory = ({ messages }) => {
  
  return (
    <div className="mx-2 px-2 border-2 border-aliceBlue-50
    rounded-xl h-48">
      <div className="p-2">
        <div className="text-white text-xl">
          Chat History
        </div>
        <div>
          {messages.map((msg, index) => (
            <div className="text-white"
            key={index}>
              {msg}
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default ChatHistory;


