# Use the file name mbox-short.txt as the file name
fname = input("Enter file name: ")
count = 0
addition_of_num = 0
fh = open(fname)
for line in fh:
    if not line.startswith("X-DSPAM-Confidence:") : continue
    count = count +1
    addition_of_num = addition_of_num+ float(line[19:])
print("Average spam confidence:",addition_of_num/count)
