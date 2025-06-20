import ChatHistory from './ChatHistory'
import Connection from './Connection'


const ChatContainer = ({messages, onEstbConn}) => {
  return (
    <div className='flex flex-col'>
      <ChatHistory messages={messages}/>
      <Connection estbConn={onEstbConn}/>
    </div>
  );
}

export default ChatContainer;