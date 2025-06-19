import goLang from './assets/goLang.svg'

function App() {
  return (
    <>
      <div className='p-2'>
        <div className='flex gap-2 items-center'>
          <img src={goLang} alt="goLang logo" className='w-8' />
          <p className='text-xl'>GoChat</p>
        </div>
      </div>
    </>
  )
}

export default App
