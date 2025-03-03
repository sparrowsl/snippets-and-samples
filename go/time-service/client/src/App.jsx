import dayjs from "dayjs";
import { createSignal } from "solid-js";

export default function App() {
	const [time, set_time] = createSignal();

	const get_current_time = async () => {
		try {
			const res = await fetch("http://localhost:5000");
			const data = await res.json();

			set_time(data.currentTime);
		} catch (_e) {
			const date = new Date();
			set_time(date.toISOString().split("T")[1].slice(0, 8));
		}
	};

	return (
		<>
			<h1>Time Service</h1>
			<div>
				<p>
					The current time is{" "}
					{dayjs(time()).format("HH:mm:ss UTC - MMMM DD, YYYY")}
				</p>

				<button type="button" onClick={get_current_time}>
					get current time
				</button>
			</div>
		</>
	);
}
