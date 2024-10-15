import { getPost } from "@/lib/api/sample";
import PostView from "./PostView";

export default async function Page({ params }: { params: { id: string } }) {
  const { id } = params;

  return (
    <div>
      <PostView id={id} />
    </div>
  );
}
