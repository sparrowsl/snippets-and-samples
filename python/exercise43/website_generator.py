import os


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


template_data = template_data.replace("{{sitename}}", site_name)
template_data = template_data.replace("{{author}}", author)

if not os.path.exists(f"./{site_name}"):
    os.mkdir(f"./{site_name}")
    print(f"Created ./{site_name}")
else:
    print(f"./{site_name} already exists")

with open(f"./{site_name}/index.html", "w") as file:
    file.write(template_data)

if js_folder.strip() == "y":
    if not os.path.exists(f"./{site_name}/js"):
        os.mkdir(f"./{site_name}/js")
        print(f"Created ./{site_name}/js")
    else:
        print(f"./{site_name}/js already exists")


if css_folder.strip() == "y":
    if not os.path.exists(f"./{site_name}/css"):
        os.mkdir(f"./{site_name}/css")
        print(f"Created ./{site_name}/css")
    else:
        print(f"./{site_name}/css already exists")
