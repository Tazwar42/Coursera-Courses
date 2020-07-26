def computepay(hour,rate):
    if hour>40:
        extra_hour = hour-40
        extra_rate = extra_hour*1.5*rate
        compute_pay = 40*rate+extra_rate
    return compute_pay

hrs = float(input("Enter Hours:"))
rate = float(input("Enter Rate:"))

print("Pay %0.2f" %computepay(hrs,rate))
