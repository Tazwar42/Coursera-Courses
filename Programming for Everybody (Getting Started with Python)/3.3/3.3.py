score = float(input("Enter Score: "))

if score>0.0 and score<1.0:
    if score>=0.9:
        print("A")
    elif score>=0.8 and score<=0.9:
        print("B")
    elif score>=0.7 and score<=0.8:
        print("C")
    elif score>=0.6 and score<=0.7:
        print("D")
    else:
        print("F")
else:
    print("You are out of the zone")