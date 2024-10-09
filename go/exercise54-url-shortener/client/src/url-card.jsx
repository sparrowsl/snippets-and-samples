/** @param {{url: {id:string, short_url:string, long_url:string}}} url  */
export default function URLCard({ url }) {
  return (
    <>
      <div class="bg-white shadow rounded p-3">
        <h3 class="text-sm">
          Short URL:{" "}
          <span class="underline text-blue-500">{url.short_url}</span>
        </h3>
        <p class="text-sm">
          Long URL:{" "}
          <span class="italic line-clamp-1 text-gray-500">{url.long_url}</span>
        </p>

        <div class="flex items-center justify-between mt-5">
          <a href={`/${url.id}/stats`} class="text-blue-700 hover:underline">
            stats &#8961;
          </a>
          <a href={url.long_url} target="_blank" class="text-blue-800 underline" rel="noreferrer">
            visit &#129146;
          </a>
        </div>
      </div>
    </>
  );
}
