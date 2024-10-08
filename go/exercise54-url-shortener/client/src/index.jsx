/* @refresh reload */
import { render } from "solid-js/web";

import "./index.css";
import { Router } from "@solidjs/router";
import { lazy } from "solid-js";

const root = document.getElementById("root");
const routes = [
	{ path: "/", component: lazy(() => import("./App.jsx")) },
	{ path: "/:id/stats", component: lazy(() => import("./url-stats.jsx")) },
];

render(() => <Router>{routes}</Router>, root);
