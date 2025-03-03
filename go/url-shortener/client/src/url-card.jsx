/** @param {{url: import("./types").Url[]}} url */
export default function URLCard({ url }) {
	/** @param {import("./types").Url} url */
	const updateVisitedThenRedirect = async (url) => {
		try {
			const res = await fetch(`http://localhost:5000/urls/${url.short_url}`, {
				method: "PATCH",
				headers: { "Content-Type": "application/json" },
				body: "",
			});
			if (res.ok) {
				window.location.href = url.long_url;
			}
		} catch (_) {}
	};

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
					<a
						href={`/${url.short_url}/stats`}
						class="text-blue-700 hover:underline"
					>
						stats &#8961;
					</a>
					<button
						class="text-blue-800 underline"
						type="button"
						onClick={() => updateVisitedThenRedirect(url)}
					>
						visit &#129146;
					</button>
				</div>
			</div>
		</>
	);
}
