import { Router } from "@solidjs/router";
import { lazy } from "solid-js";
import { render } from "solid-js/web";

import "./index.css";

const root = document.getElementById("root");
const routes = [
	{ path: "/", component: lazy(() => import("./App.jsx")) },
	{ path: "/:id", component: lazy(() => import("./redirect.jsx")) },
	{ path: "/:id/stats", component: lazy(() => import("./url-stats.jsx")) },
];

render(() => <Router>{routes}</Router>, root);
