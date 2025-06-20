import goLang from "../assets/goLang.svg"

const Navbar = () => {
  return (
    <div className='p-2 flex justify-center'>
      <div className='flex gap-1 items-center rounded-xl bg-amber-50 w-32 p-2 justify-center shadow-xl'>
        <img src={goLang} alt="goLang logo" className='w-7 -ml-1' />
        <p className='text-xl'>GoChat</p>
      </div>
    </div>
  );
}

export default Navbar;