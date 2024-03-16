/*
interface Props{
	title: string;
	likes: number;
	caption: string;
	image: string;
}
*/

export default function Post(/*{title,likes,caption,image}:Props*/) {
  return (
    <>
	  <div className="flex justify-center font-bold text-2xl">Title</div>
	  <hr className="h-px my-1 bg-gray-300 border-0"/>
	  
    </>
  );
}
