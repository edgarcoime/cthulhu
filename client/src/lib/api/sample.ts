import { API_BASE_URL } from "@/constants/api";

export interface Post {
  id: string;
  title: string;
  body: string;
  userId: number;
}

export async function getPosts(): Promise<Post[] | null> {
  try {
    const url = `${API_BASE_URL}/posts`;
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

export async function getPost(id: number): Promise<Post | null> {
  const url = `${API_BASE_URL}/posts/${id}`;
  const res = await fetch(url, { cache: "no-store" });

  if (!res.ok) {
    //throw new Error(`Error fetching post ${id}`);
  }

  const post = await res.json();
  return post;
}
