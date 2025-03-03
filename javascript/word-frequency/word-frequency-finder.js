import fs from "node:fs";
import { serve } from "@hono/node-server";
import { Hono } from "hono";

const app = new Hono();

const contents = fs.readFileSync("./macbeth.txt", "utf8").trim();

/** @type {Map<string, number>} store */
const store = new Map();

const splitted = contents.split(/ |\n/);

for (const word of splitted) {
	const trimmed = word.replace(/[,;:_?!@\"'.—”“’\)\(&*%$#-|I],/g, "");

	if (trimmed === "") {
		continue;
	}

	store.set(trimmed, store.get(trimmed) + 1 || 1);
}

console.log(Object.fromEntries(store));

app.get("/", (c) => {
	return c.render("<h1>Hello word finder!!</h1>");
});

serve({ fetch: app.fetch, port: 5000 });
