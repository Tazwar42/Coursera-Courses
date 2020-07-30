import re
sum = 0
name = input("Enter file:")
if len(name) < 1 : name = "regex_sum_836430.txt"
handle = open(name)
# read the whole text
text = handle.read()
stuff = re.findall('[0-9]+',text)
# loop through stuff list and summation
for num in stuff:
    sum = sum + int(num)
print(sum)
