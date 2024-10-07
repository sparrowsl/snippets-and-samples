import { createSignal } from "solid-js";

function App() {
	const [count, setCount] = createSignal(0);

	return (
		<>
			<section class="">
				<h1>Breve - URL Shortner</h1>

				<div>
					<button type="button" onClick={() => setCount((count) => count + 1)}>
						count is {count()}
					</button>
				</div>
			</section>
		</>
	);
}

export default App;
