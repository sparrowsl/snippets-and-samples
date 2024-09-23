import fs from "node:fs";
import { input } from "pynput";

const template_data = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="author" content="{{author}}">
    <title>{{sitename}}</title>
  </head>
  <body>
  </body>
</html>`;

const site_name = await input("Site name: ");
const author = await input("Author: ");
const js_folder = await input("Do you want a folder for JavaScript? ");
const css_folder = await input("Do you want a folder for CSS? ");

if (!fs.existsSync(`./${site_name}`)) {
	fs.mkdirSync(site_name);
}
console.log(`Created: ./${site_name}`);

const replaced = template_data
	.replace(/{{sitename}}/g, site_name)
	.replace(/{{author}}/g, author);

fs.writeFileSync(`./${site_name}/index.html`, replaced);
console.log(`Created: ./${site_name}/index.html`);

if (js_folder === "y") {
	if (!fs.existsSync(`./${site_name}/js`)) {
		fs.mkdirSync(`./${site_name}/js/`, { recusive: true });
	}

	console.log(`Created: ./${site_name}/js/`);
}

if (css_folder === "y") {
	if (!fs.existsSync(`./${site_name}/css`)) {
		fs.mkdirSync(`./${site_name}/css/`, { recusive: true });
	}

	console.log(`Created: ./${site_name}/css/`);
}
