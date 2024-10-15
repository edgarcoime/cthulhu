import NavButton from "@/components/NavButton";

export default async function Home() {
  return (
    <div>
      <h1>Cthulhu Microservices App</h1>
      <NavButton path="/posts" desc="Posts" />
    </div>
  );
}
