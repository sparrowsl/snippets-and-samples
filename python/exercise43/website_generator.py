from pathlib import Path

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


site_name = input("Site name: ")
author = input("Author: ")
js_folder = input("Do you want a folder for JavaScript? ")
css_folder = input("Do you want a folder for CSS? ")


Path(f"./{site_name}").mkdir(parents=True, exist_ok=True)
with open(f"./{site_name}/index.html", "w") as file:
    template_data = template_data.replace("{{sitename}}", site_name)
    template_data = template_data.replace("{{author}}", author)

    file.write(template_data)
    print(f"Created ./{site_name}")
    print(f"Created ./{site_name}/index.html")

if js_folder.strip() == "y":
    Path(f"./{site_name}/js").mkdir(parents=True, exist_ok=True)
    print(f"Created ./{site_name}/js")

if css_folder.strip() == "y":
    Path(f"./{site_name}/css").mkdir(parents=True, exist_ok=True)
    print(f"Created ./{site_name}/css")
