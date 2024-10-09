import { A, useParams } from "@solidjs/router";
import { createResource } from "solid-js";

export default function URLStats() {
  const params = useParams();
  const [data] = createResource(async () => {
    const res = await fetch(`http://localhost:5000/urls/${params.id}/stats`);
    const data = await res.json();
    return data.url;
  });

  return (
    <>
      <section class="relative max-w-5xl min-h-screen mx-auto bg-gray-50 py-5">
        <h1 class="text-xl font-bold text-center underline">
          Breve - URL Shortner
        </h1>
        <A
          href="/"
          class="block w-fit text-sm mt-2 mx-auto underline text-blue-400"
        >
          Home
        </A>

        <div class="bg-white rounded p-10 max-w-96 mx-auto mt-10 shadow">
          <h2 class="font-semibold">
            Short URL:{" "}
            <span class="underline text-blue-900 ml-4 font-normal">
              {data()?.short_url}
            </span>{" "}
          </h2>
          <p class="font-semibold">
            Long URL:{" "}
            <span class="text-sm font-normal ml-4">{data()?.long_url}</span>
          </p>
          <p class="font-semibold">
            Visited:{" "}
            <span class="text-sm font-normal ml-4">{data()?.visited}</span>
          </p>
        </div>
      </section>
    </>
  );
}
