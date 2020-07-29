name = input("Enter file:")
if len(name) < 1 : name = "mbox-short.txt"
handle = open(name)
counts = dict()
ls = list()

# refine the hour from the file to a list
for line in handle:
    if line.startswith('From '):
        words = line.split()
        #words = line.split(':')
        ls.append(words[5].split(':')[0])
#print(ls)

# update/create dict with keys and values
for word in ls:
    counts[word] = counts.get(word,0) + 1
#print(counts)

#print(sorted([(k,v) for k,v in counts.items()]))
ls2 = list()
# make a tuple and append the value in list and sort it
for key,val in counts.items():
    newtup = (key,val)
    ls2.append(newtup)
ls2 = sorted(ls2)
for v,k in ls2:
    print(v,k)
