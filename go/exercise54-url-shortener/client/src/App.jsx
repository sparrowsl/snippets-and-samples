import { createResource, For, Show } from "solid-js";

export default function App() {
	const get_urls = async () => {
		const res = await fetch("http://localhost:5000/urls");
		const data = await res.json();
		return data.urls
	};

	const [data] = createResource(get_urls)



	return (
		<>
			<section class="max-w-5xl min-h-screen mx-auto bg-gray-50 py-5">
				<h1 class="text-xl font-bold text-center underline">
					Breve - URL Shortner
				</h1>

				<Show when={data.loading}>
					<p>loading urls....</p>
				</Show>

				<div class="grid grid-cols-3 gap-5 p-10">
					<For each={data()} fallback={<p>No urls yet!!</p>}>
						{(url) => <Card url={url} />}
					</For>
				</div>
			</section>
		</>
	);
}

/** @param {{url: {id:string, short_url:string, long_url:string}}} url  */
function Card({ url }) {
	return (
		<>
			<div class="bg-white shadow rounded p-3">
				<h3>Short URL: &copy; {url.short_url}</h3>
				<p class="text-sm">Long URL: &trade; {url.long_url}</p>

				<a href="/" class="text-blue-800 block w-fit underline mt-5">
					visit &rarr;
				</a>
			</div>
		</>
	);
}
