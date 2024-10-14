export interface Post {
  id: string;
  title: string;
  body: string;
  userId: number;
}

export async function getPosts(): Promise<Post[] | null> {
  const url = "https://jsonplaceholder.typicode.com/posts";

  try {
    const res = await fetch(url);

    if (!res.ok) {
      console.log("Error fetching posts.");
      return null;
    }

    const posts = await res.json();
    return posts;
  } catch (error) {
    console.log(error);
    return null;
  }
}
