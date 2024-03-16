import Post from "../components/Post"
function App() {

  return (
    <>
	  <div className="flex justify-center">
	  <input className="bg-gray-200 rounded-lg p-1 px-5 my-3" placeholder="Search..." />
	  </div>
	  <hr className="h-px my-1 bg-gray-300 border-0"/>


	  <Post/>
    </>
  )
}

export default App;
