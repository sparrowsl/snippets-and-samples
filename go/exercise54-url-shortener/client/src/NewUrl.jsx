import { createSignal } from "solid-js";

export default function NewUrl() {
	const [url, set_url] = createSignal();

	const add_url = async (e) => {
		e.preventDefault();

		if (url().trim() === "") {
			return;
		}

		await fetch("http://localhost:5000/urls", {
			method: "POST",
			headers: { "Content-Type": "application/json" },
			body: JSON.stringify({ url: url() }),
		});
	};

	return (
		<>
			<section>
				<form action="" method="POST" onSubmit={add_url}>
					<input
						type="text"
						name="long_url"
						placeholder="https://www.google.com"
						class="block text-sm rounded p-1"
						onChange={(e) => set_url(e.target.value)}
					/>
					<button
						type="submit"
						class="block bg-gray-300 text-sm rounded max-w-fit py-2 px-6"
					>
						Submit
					</button>
				</form>
			</section>
		</>
	);
}
