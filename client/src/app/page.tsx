import { getPosts, Post } from "@/lib/api/sample";

function PostsView({ posts }: { posts: Post[] }) {
  return (
    <>
      {posts.map((p, idx) => (
        <div key={idx}>
          <h2>{p.title}</h2>
          <p>{p.body}</p>
        </div>
      ))}
    </>
  );
}

export default async function Home() {
  const posts = await getPosts();

  return (
    <div>
      <h1>Cthulhu Microservices App</h1>
      {posts ? <PostsView posts={posts} /> : <p>Loading...</p>}
    </div>
  );
}
