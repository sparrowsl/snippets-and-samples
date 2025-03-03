import { For, Show, createResource, createSignal } from "solid-js";

import NewUrl from "./new-url.jsx";
import URLCard from "./url-card.jsx";

export const [open, set_open] = createSignal(false);

export default function App() {
	const get_urls = async () => {
		try {
			const res = await fetch("http://localhost:5000/urls");
			/** @type {{urls: import("./types").Url[]}} */
			const data = await res.json();

			return data.urls;
		} catch (_e) {
			return [];
		}
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
						{(url) => <URLCard url={url} />}
					</For>
				</div>
			</section>
		</>
	);
}
