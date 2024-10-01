import dayjs from "dayjs";
import { createSignal } from "solid-js";

export default function App() {
	const [time, set_time] = createSignal();

	const get_current_time = async () => {
		const res = await fetch("http://localhost:5000");
		const data = await res.json();

		set_time(data.currentTime);
	};

	return (
		<>
			<h1>Time Service</h1>
			<div class="card">
				<button type="button" onClick={get_current_time}>
					get current time
				</button>

				<p>
					The current time is {dayjs(time()).format("HH:mm:ss UTC MMMM d YYYY")}
				</p>
			</div>
		</>
	);
}
