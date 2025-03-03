import { createSignal } from "solid-js";

export default function NewUrl() {
	const [url, set_url] = createSignal();

	const add_url = async (e) => {
		e.preventDefault();

		if (url().trim() === "") {
			return;
		}

		try {
			const res = await fetch("http://localhost:5000/urls", {
				method: "POST",
				headers: { "Content-Type": "application/json" },
				body: JSON.stringify({ url: url() }),
			});

			if (!res.ok) {
				return;
			}
		} catch (_) {}

		window.location.reload();
	};

	return (
		<>
			<section>
				<form method="POST" onSubmit={add_url}>
					<legend class="text-center font-semibold text-sm">Add New URL</legend>

					<input
						type="text"
						name="long_url"
						placeholder="https://www.google.com"
						class="block text-sm rounded p-2 w-full my-5"
						onChange={(e) => set_url(e.target.value)}
					/>
					<button
						type="submit"
						class="cursor-pointer mx-auto hover:bg-blue-400 block bg-gray-300 text-sm rounded max-w-fit py-2 px-6"
					>
						Submit
					</button>
				</form>
			</section>
		</>
	);
}
