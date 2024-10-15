import { getPosts, Post } from "@/lib/api/sample";
import PostsView from "./PostsView";

export default async function Home() {
  const posts = await getPosts();

  return (
    <div>
      <h1>Cthulhu Microservices App</h1>
      <section>
        {posts ? <PostsView posts={posts} /> : <p>Loading...</p>}
      </section>
    </div>
  );
}
