"use client";

import { Post } from "@/lib/api/sample";
import { useRouter } from "next/navigation";

function PostTile({ post }: { post: Post }) {
  const url = `/posts/${post.id}`;
  const router = useRouter();

  function clickHandler() {
    router.push(url);
  }

  return (
    <div onClick={clickHandler} className="bg-gray-900 my-5 cursor-pointer">
      <h2>{post.title}</h2>
      <p>{post.body}</p>
    </div>
  );
}

export default function PostsView({ posts }: { posts: Post[] }) {
  return (
    <div>
      {posts.map((post, _) => (
        <PostTile key={post.id} post={post} />
      ))}
    </div>
  );
}
