import { getPost } from "@/lib/api/sample";
import { ErrorBoundary } from "react-error-boundary";
import { Suspense } from "react";

function ErrorView() {
  return <div>Something went wrong</div>;
}

function LoadingView() {
  return <div>loading...</div>;
}

async function FetchDataView({ id }: { id: string }) {
  const post = await getPost(Number(id));
  return (
    <div>
      <h2>Post {id}</h2>
      <pre>{JSON.stringify(post, null, 2)}</pre>
    </div>
  );
}

export default function PostView({ id }: { id: string }) {
  return (
    <div>
      <Suspense fallback={<LoadingView />}>
        <ErrorBoundary fallback={<ErrorView />}>
          <FetchDataView id={id} />
        </ErrorBoundary>
      </Suspense>
    </div>
  );
}
