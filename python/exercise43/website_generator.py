import pathlib
import zipfile

template_data = """
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
</html>
"""


def zip_dir(filename: str):
    current_path = pathlib.Path(filename)

    with zipfile.ZipFile(f"{filename}.zip", "w", zipfile.ZIP_DEFLATED) as zip_file:
        for entry in current_path.rglob("*"):
            zip_file.write(entry, entry.relative_to(current_path))


def create_dir(pathname: str):
    pathlib.Path(f"./{pathname}").mkdir(parents=True, exist_ok=True)
    print(f"Created ./{pathname}")


site_name = input("Site name: ").strip()
author = input("Author: ").strip()
js_folder = input("Do you want a folder for JavaScript? ").strip()
css_folder = input("Do you want a folder for CSS? ").strip()

create_dir(site_name)
with open(f"./{site_name}/index.html", "w") as file:
    template_data = template_data.replace("{{sitename}}", site_name)
    template_data = template_data.replace("{{author}}", author)

    file.write(template_data)
    print(f"Created ./{site_name}/index.html")

if js_folder == "y":
    create_dir(f"{site_name}/js")

if css_folder == "y":
    create_dir(f"{site_name}/css")

zip_dir(site_name)
