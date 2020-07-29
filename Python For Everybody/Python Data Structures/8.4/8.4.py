fname = input("Enter file name: ")
fh = open(fname)
list_1 = list()
list_2 = list()
for line in fh:
    list_2 = line.split()
    for word in list_2:
        if word in list_1:
            continue
        list_1.append(word)
list_1.sort()
print(list_1)
