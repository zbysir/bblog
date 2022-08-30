import Link from "./Link";
import {BlogI} from "./BlogSmall";


export default function BlogXS({blog}: { blog: BlogI }) {
    let link = '/blogs/' + blog.name
    let name = blog.meta?.title || blog.name

    return <div className="">
        {/*{*/}
        {/*    blog.meta?.img ? <Link href={link} className="block w-20 md:w-60 md:h-20">*/}
        {/*        <img*/}
        {/*            className="object-cover mb-2 overflow-hidden rounded-lg shadow-sm w-full h-full"*/}
        {/*            src={blog.meta?.img}/>*/}
        {/*    </Link> : null*/}
        {/*}*/}
        <div className="flex items-center space-x-4">
            <div className="w-1/2 text-right">
                <h2 className="font-bold text-xl">
                    <Link href={link}> {name}</Link>
                </h2>

            </div>
            <div className="w-1/2 flex flex-col space-y-1">
                <p className="text-sm text-gray-500">{blog.meta.desc}</p>

                <div className="flex space-x-3">
                    {
                        (function () {
                            let tags = blog.meta?.tags
                            if (typeof tags === "string") {
                                tags = [tags]
                            }

                            return tags?.map(i => (
                                <Link href={"/tags/" + i}>
                                    <div
                                        className="bg-gray-500 items-center px-1 py-0.5 leading-none rounded-full text-xs font-medium text-white ">
                                        <span>{i}</span>
                                    </div>
                                </Link>
                            ))
                        })()
                    }
                </div>
            </div>


        </div>


    </div>

}