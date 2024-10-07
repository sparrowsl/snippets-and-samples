import { For } from "solid-js";

export default function App() {
	return (
		<>
			<section class="max-w-5xl min-h-screen mx-auto bg-gray-50 py-5">
				<h1 class="text-xl font-bold text-center underline">
					Breve - URL Shortner
				</h1>

				<div class="grid grid-cols-3 gap-5 p-10">
					<For each={Array.from({ length: 5 })}>
						{<Card url={{ id: "", short_url: "short", long_url: "" }} />}
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
				<p>Long URL: &trade; {url.long_url}</p>

				<a href="/" class="text-blue-800 block w-fit underline mt-5">
					visit &rarr;
				</a>
			</div>
		</>
	);
}
