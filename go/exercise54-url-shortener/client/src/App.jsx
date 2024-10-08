import { createResource, createSignal, For, Show } from "solid-js";
import NewUrl from "./NewUrl";

export default function App() {
	const [open, set_open] = createSignal(false);

	const get_urls = async () => {
		const res = await fetch("http://localhost:5000/urls");
		const data = await res.json();
		return data.urls;
	};

	const [data] = createResource(get_urls);

	return (
		<>
			<section class="relative max-w-5xl min-h-screen mx-auto bg-gray-50 py-5">
				<h1 class="text-xl font-bold text-center underline">
					Breve - URL Shortner
				</h1>
				<button
					type="button"
					class="block w-fit text-xs mx-auto bg-blue-400 font-semibold px-4 py-2 cursor-pointer text-white rounded"
					id="dialog"
					onClick={() => set_open(true)}
				>
					Add URL
				</button>

				<Show when={open()}>
					<div class="absolute inset-0 grid place-content-center bg-black/20">
						<div class="bg-white p-10 rounded shadow relative min-w-96">
							<button
								type="button"
								class="absolute top-1 right-3 text-2xl text-red-400 cursor-pointer"
								onClick={() => set_open(false)}
							>
								&times;
							</button>
							<NewUrl />
						</div>
					</div>
				</Show>

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
				<h3 class="text-sm">
					Short URL:{" "}
					<span class="underline text-blue-500">{url.short_url}</span>
				</h3>
				<p class="text-sm">
					Long URL:{" "}
					<span class="italic line-clamp-1 text-gray-500">{url.long_url}</span>
				</p>

				<div class="flex items-center justify-between mt-5">
					<a href="/" class="text-blue-700 hover:underline">
						stats &#8961;
					</a>
					<a href="/" class="text-blue-800 underline">
						visit &#129146;
					</a>
				</div>
			</div>
		</>
	);
}
