"use client";

import { useRouter } from "next/navigation";

export default function NavButton({
  path,
  desc,
}: {
  path: string;
  desc: string;
}) {
  const router = useRouter();

  function clickHandler() {
    router.push(path);
  }

  return <button onClick={clickHandler}>{desc}</button>;
}
