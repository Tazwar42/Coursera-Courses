hrs = float(input("Enter Hours:"))
rate_per_hour = float(input("Enter Rate Per Hour:"))
limit = 40

if(hrs>40):
    extra_hour = hrs - limit
    rate_per_extra_hour = extra_hour*1.5
    print((limit*rate_per_hour)+(rate_per_hour*rate_per_extra_hour))
else:
    print(hrs*rate_per_hour)
