length = int(input("What is the length of room in feet? "))
width = int(input("What is the width of room in feet? "))
square_meters = (float(length) * float(width) * 0.09290304)

print(f"You entered dimensions of {length} feet by {width} feet.")
print(f"The area is\n{length * width} square feet\n{square_meters:.3f} square meters")
