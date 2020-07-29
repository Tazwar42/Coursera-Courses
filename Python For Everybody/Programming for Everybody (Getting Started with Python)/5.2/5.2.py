largest = 0
smallest = None
while True:
    num = input("Enter a number: ")
    if num == "done" : break
    try:
        int_num = int(num)
        if largest<int_num:
            largest = int_num
        if smallest is None:
            smallest = int_num
        if smallest>int_num:
            smallest = int_num
    except:
        print("Invalid input")
print("Maximum is", largest)
print("Minimum is", smallest)
